package article

import (
	"context"

	"github.com/whuangz/micro-blog/article/proto/article"
)

type Usecase interface {
	FetchArticles(ctx context.Context, req *article.ListArticleRequest) ([]*article.Article, error)
	CreateArticle(ctx context.Context, req *article.CreateArticleRequest) (int32, error)
	DeleteArticle(ctx context.Context, req *article.DeleteArticleRequest) (int32, error)
	UpdateArticle(ctx context.Context, req *article.UpdateArticleRequest) error
}

type Repository interface {
	FetchArticles(ctx context.Context, req *article.ListArticleRequest) ([]*article.Article, error)
	CreateArticle(ctx context.Context, req *article.CreateArticleRequest) error
	DeleteArticle(ctx context.Context, req *article.DeleteArticleRequest) error
	UpdateArticle(ctx context.Context, req *article.UpdateArticleRequest) error
}
