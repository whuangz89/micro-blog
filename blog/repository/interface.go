package repository

import (
	"context"

	pb "github.com/whuangz/micro-blog/blog/proto"
)

type Service interface {
	FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error)
	CreateArticle(ctx context.Context, req *pb.Article) error
	DeleteArticle(ctx context.Context, req *pb.Article) error
	UpdateArticle(ctx context.Context, req *pb.Article) error

	FetchCategories(ctx context.Context, req *pb.ListCategoryRequest) ([]*pb.Category, error)
	CreateCategory(ctx context.Context, req *pb.Category) error

	Ping() error
	Close() error
}
