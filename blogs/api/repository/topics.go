package repository

import (
	"context"

	"github.com/gosimple/slug"
	"github.com/sirupsen/logrus"

	pb "github.com/whuangz/micro-blog/blogs/proto"
)

func (m *articleRepository) FetchTopics(ctx context.Context, req *pb.ListRequest) ([]*Topic, error) {

	query := `SELECT id, title, slug, tag, color, updated_at, created_at
						FROM topics`

	rows, err := m.db.Raw(query).Rows()
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*Topic, 0)
	for rows.Next() {
		t := new(Topic)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Slug,
			&t.Tag,
			&t.Color,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *articleRepository) CreateTopic(ctx context.Context, req *pb.Topic) error {
	query := `INSERT INTO topics(title, slug, tag, color, updated_at, created_at) VALUES(?,?,?,?,"drafted",NOW(),NOW())`
	q := m.db.Exec(query, req.Title, slug.Make(req.Title), req.Tag, req.Color)
	if errs := q.GetErrors(); len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func (m *articleRepository) GetTopic(ctx context.Context, req *pb.Topic) (*Topic, error) {
	query := `SELECT id, title, slug, tag, color FROM topics WHERE id = ?`
	rows, err := m.db.Raw(query, req.Id).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	t := new(Topic)
	for rows.Next() {
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Slug,
			&t.Tag,
			&t.Color,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
	}

	return t, nil
}
