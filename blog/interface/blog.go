package article

import (
	"context"

	pb "github.com/whuangz/micro-blog/blog/proto/blog"
)

type Usecase interface {
	FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error)
	CreateArticle(ctx context.Context, req *pb.Article) (int32, error)
	DeleteArticle(ctx context.Context, req *pb.Article) (int32, error)
	UpdateArticle(ctx context.Context, req *pb.Article) error
}

type Repository interface {
	FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error)
	CreateArticle(ctx context.Context, req *pb.Article) error
	DeleteArticle(ctx context.Context, req *pb.Article) error
	UpdateArticle(ctx context.Context, req *pb.Article) error
}
