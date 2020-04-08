package handler

import (
	"context"
	"database/sql"

	"github.com/micro/go-micro/v2/errors"
	article "github.com/whuangz/micro-blog/article/interface"
	pb "github.com/whuangz/micro-blog/article/proto/article"
	"github.com/whuangz/micro-blog/article/repository"
	"github.com/whuangz/micro-blog/article/usecase"
)

type handler struct {
	usecase article.Usecase
}

func NewArticleHandler(dbConn *sql.DB) *handler {

	articleRepo := repository.NewMysqlArticleRepository(dbConn)
	au := usecase.NewArticleUsecase(articleRepo)

	return &handler{
		usecase: au,
	}
}

func (h *handler) FetchArticles(ctx context.Context, req *pb.ListArticleRequest, res *pb.Result) error {
	articles, err := h.usecase.FetchArticles(ctx, req)
	if err != nil {
		return err
	}

	res.Articles = articles
	return nil
}

func (h *handler) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest, res *pb.Result) error {

	if req.Title == "" || req.Content == "" {
		return errors.BadRequest("", "Missing Param")
	}

	status, err := h.usecase.CreateArticle(ctx, req)
	if err != nil {
		return err
	}
	res.StatusCode = status
	res.Message = "Success Inserted Article"
	return nil
}

func (h *handler) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest, res *pb.Result) error {

	if req.Id == 0 {
		return errors.BadRequest("", "Missing Article ID")
	}

	status, err := h.usecase.DeleteArticle(ctx, req)
	if err != nil {
		return err
	}
	res.StatusCode = status
	res.Message = "Successfully deleted Article"
	return nil
}
