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

func (a *articleUsecase) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (int32, error) {
	err := a.articleRepo.CreateArticle(ctx, req)
	if err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *articleUsecase) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (int32, error) {
	err := a.articleRepo.DeleteArticle(ctx, req)
	if err != nil {
		return 400, err
	}
	return 200, nil
}
