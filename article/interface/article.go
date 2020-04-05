package article

import (
	"context"

	"github.com/whuangz/micro-blog/article/proto/article"
)

type Usecase interface {
	FetchArticles(ctx context.Context, req *article.ListArticleRequest) ([]*article.Article, error)
}

type Repository interface {
	FetchArticles(ctx context.Context, req *article.ListArticleRequest) ([]*article.Article, error)
}
