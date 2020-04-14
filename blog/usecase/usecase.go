package usecase

import (
	"context"

	blog "github.com/whuangz/micro-blog/blog/interface"
	pb "github.com/whuangz/micro-blog/blog/proto/blog"
)

type blogUsecase struct {
	blogRepo blog.Repository
}

func NewBlogUsecase(b blog.Repository) blog.Usecase {
	return &blogUsecase{
		blogRepo: b,
	}
}

func (a *blogUsecase) FetchArticles(ctx context.Context, req *pb.ListArticleRequest) ([]*pb.Article, error) {
	listArticle, err := a.blogRepo.FetchArticles(ctx, req)
	if err != nil {
		return nil, err
	}

	return listArticle, nil
}

func (a *blogUsecase) CreateArticle(ctx context.Context, req *pb.Article) (int32, error) {
	err := a.blogRepo.CreateArticle(ctx, req)
	if err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *blogUsecase) DeleteArticle(ctx context.Context, req *pb.Article) (int32, error) {
	err := a.blogRepo.DeleteArticle(ctx, req)
	if err != nil {
		return 400, err
	}
	return 200, nil
}

func (a *blogUsecase) UpdateArticle(ctx context.Context, req *pb.Article) error {
	err := a.blogRepo.UpdateArticle(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
