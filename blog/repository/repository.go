package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"

	"github.com/micro/go-micro/v2/errors"
	"github.com/sirupsen/logrus"
	pb "github.com/whuangz/micro-blog/blog/proto"
)

type blogRepository struct {
	Conn *sql.DB
}

func New(dbHost, dbName, dbUser, dbPass, dbPort string) (Service, error) {

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		return nil, err
	}

	return &blogRepository{dbConn}, nil
}

// Close will terminate the database connection
func (m *blogRepository) Close() error {
	return m.Conn.Close()
}

// Close will terminate the database connection
func (m *blogRepository) Ping() error {
	return m.Conn.Ping()
}

func (m *blogRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*pb.Article, error) {
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

func (m *blogRepository) FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article`

	res, err := m.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *blogRepository) CreateArticle(ctx context.Context, req *pb.Article) error {
	query := `INSERT INTO article(title, content, author_id, updated_at, created_at) VALUES(?,?,?,NOW(),NOW())`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, req.Title, req.Content, 1)
	if err != nil {
		return err
	}

	return nil
}

func (m *blogRepository) DeleteArticle(ctx context.Context, req *pb.Article) error {
	query := `DELETE FROM article WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, req.Id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected != 1 {
		return errors.NotFound("", "Cannot Delete Article with this Id")
	}

	return nil
}

func (m *blogRepository) UpdateArticle(ctx context.Context, req *pb.Article) error {
	query := `UPDATE article SET title=?, content=?, updated_at=NOW() WHERE id=?`
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, req.Title, req.Content, req.Id)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if rowsAfected != 1 {
		return errors.NotFound("", "Article Not Found")
	}

	return nil
}
