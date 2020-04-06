package article

import (
	"context"

	"github.com/whuangz/micro-blog/article/proto/article"
)

type Usecase interface {
	FetchArticles(ctx context.Context, req *article.ListArticleRequest) ([]*article.Article, error)
	CreateArticle(ctx context.Context, req *article.CreateArticleRequest) (int32, error)
}

type Repository interface {
	FetchArticles(ctx context.Context, req *article.ListArticleRequest) ([]*article.Article, error)
	CreateArticle(ctx context.Context, req *article.CreateArticleRequest) error
}
