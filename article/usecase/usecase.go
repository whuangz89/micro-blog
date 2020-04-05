package usecase

import (
	"context"

	"github.com/whuangz/micro-blog/article/interface"
	pb "github.com/whuangz/micro-blog/article/proto/article"
)

type articleUsecase struct {
	articleRepo article.Repository
}

func NewArticleUsecase(a article.Repository) article.Usecase {
	return &articleUsecase{
		articleRepo: a,
	}
}

func (a *articleUsecase) FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error) {
	listArticle, err := a.articleRepo.FetchArticles(ctx, req)
	if err != nil {
		return nil, err
	}

	return listArticle, nil
}
