package handler

import (
	"context"
	"database/sql"

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
