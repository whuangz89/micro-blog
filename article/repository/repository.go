package repository

import (
	"context"
	"database/sql"

	"github.com/sirupsen/logrus"
	"github.com/whuangz/micro-blog/article/interface"
	pb "github.com/whuangz/micro-blog/article/proto/article"
)

type articleRepositroy struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) article.Repository {
	return &articleRepositroy{Conn}
}

func (m *articleRepositroy) fetch(ctx context.Context, query string, args ...interface{}) ([]*pb.Article, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
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
		t := new(pb.Article)
		authorID := int64(0)
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		t.Author = &pb.Author{
			Id: authorID,
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *articleRepositroy) FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article`

	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
