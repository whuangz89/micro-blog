package repository

import (
	"context"
	"strings"

	"github.com/gosimple/slug"
	"github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"

	pb "github.com/whuangz/micro-blog/blogs/proto"
)

func (m *articleRepository) fetch(query string) ([]*Article, error) {

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

	result := make([]*Article, 0)
	for rows.Next() {
		t := new(Article)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Slug,
			&t.Content,
			&t.TopicID,
			&t.AuthorID,
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

func (m *articleRepository) FetchArticles() ([]*Article, error) {

	query := `SELECT id,title, slug, content, topic_id, author_id, updated_at, created_at
						FROM articles`

	res, err := m.fetch(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *articleRepository) CreateArticle(ctx context.Context, req *pb.Article) error {
	query := `INSERT INTO articles(title, slug, content ,topic_id ,author_id, status ,updated_at, created_at) VALUES(?,?,?,?,?,"drafted",NOW(),NOW())`
	q := m.db.Exec(query, strings.TrimSpace(req.Title), slug.Make(req.Title), strings.TrimSpace(req.Content), 1, 1)
	if errs := q.GetErrors(); len(errs) > 0 {
		return errs[0]
	}

	return nil
}

func (m *articleRepository) DeleteArticle(ctx context.Context, req *pb.Article) error {
	query := `DELETE FROM articles where id=?`
	q := m.db.Exec(query, req.Id)
	if errs := q.GetErrors(); len(errs) > 0 {
		return errs[0]
	}

	rowAffected := q.RowsAffected
	if rowAffected < 1 {
		return errors.NotFound("", "Cannot Delete Article with this Id")
	}

	return nil
}

func (m *articleRepository) UpdateArticle(ctx context.Context, req *pb.Article) error {
	// query := `UPDATE articles SET id,title, slug, content, category_id, author_id, updated_at, created_at
	// 					FROM articles`
	return nil
}
