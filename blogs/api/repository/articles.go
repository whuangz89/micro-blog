package repository

import (
	"context"
	"strings"

	"github.com/gosimple/slug"
	"github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"

	pb "github.com/whuangz/micro-blog/blogs/proto"
	"github.com/whuangz/micro-blog/helpers/sql"
)

func (m *articleRepository) fetch(query string) ([]*pb.Article, error) {

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

	result := make([]*pb.Article, 0)
	for rows.Next() {
		a := new(Article)
		t := new(Topic)
		au := new(Author)

		err = rows.Scan(
			&a.ID,
			&a.Title,
			&a.Slug,
			&a.Content,
			&t.ID,
			&t.Title,
			&t.Slug,
			&t.Color,
			&au.ID,
			&au.Name,
			&a.IsPublished,
			&a.UpdatedAt,
			&a.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		result = append(result, m.serializeArticle(a, t, au))
	}

	return result, nil

}

func (m *articleRepository) FetchArticles(req *pb.ListRequest) ([]*pb.Article, error) {

	status := req.Status.String()
	var whereQuery = `WHERE a.deleted_at is NULL`
	if status == "DRAFTED" {
		whereQuery += ` AND is_published = 0`
	} else {
		whereQuery += ` AND is_published = 1`
	}

	query := `
			SELECT a.id, a.title, a.slug, a.content, 
			a.topic_id , t.title as topic_title, t.slug as topic_slug, t.color as color_code,
			a.author_id, au.name as author_name,
			a.is_published, a.updated_at, a.created_at
			FROM articles a
			LEFT JOIN authors au ON au.id = a.author_id AND au.deleted_at is NULL
			LEFT JOIN topics t ON t.id = a.topic_id AND t.deleted_at is NULL
			`
	query += whereQuery

	res, err := m.fetch(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *articleRepository) CreateArticle(ctx context.Context, req *pb.Article) error {
	query := `
			INSERT INTO articles(title, slug, content ,topic_id ,author_id, is_published, published_at ,updated_at, created_at) 
			VALUES(?,?,?,?,?,?,?,NOW(),NOW())
			`

	q := m.db.Exec(query,
		strings.TrimSpace(req.Title), slug.Make(req.Title), strings.TrimSpace(req.Content), req.TopicId, 1, req.PublishedAt != "", sql.NewNullString(req.PublishedAt))

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

func (h *articleRepository) serializeArticle(a *Article, t *Topic, au *Author) *pb.Article {
	return &pb.Article{
		Id:      a.ID,
		Title:   a.Title,
		Slug:    a.Slug,
		Content: a.Content,
		Topic: &pb.Topic{
			Id:    t.ID,
			Title: t.Title,
			Slug:  t.Slug,
			Color: t.Color,
		},
		Author: &pb.Author{
			Id:   au.ID,
			Name: au.Name,
		},
		IsPublished: a.IsPublished,
		CreatedAt:   a.CreatedAt.Unix(),
		UpdatedAt:   a.UpdatedAt.Unix(),
	}
}
