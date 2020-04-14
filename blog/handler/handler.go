package handler

import (
	"context"
	"database/sql"

	"github.com/micro/go-micro/v2/errors"
	article "github.com/whuangz/micro-blog/blog/interface"
	pb "github.com/whuangz/micro-blog/blog/proto/blog"
	"github.com/whuangz/micro-blog/blog/repository"
	"github.com/whuangz/micro-blog/blog/usecase"
)

type handler struct {
	usecase article.Usecase
}

func NewArticleHandler(dbConn *sql.DB) *handler {

	blogRepo := repository.NewMysqlBlogRepository(dbConn)
	au := usecase.NewBlogUsecase(blogRepo)

	return &handler{
		usecase: au,
	}
}

func (h *handler) FetchArticles(ctx context.Context, req *pb.ListArticleRequest, res *pb.ArticleResult) error {
	articles, err := h.usecase.FetchArticles(ctx, req)
	if err != nil {
		return err
	}

	res.Articles = articles
	return nil
}

func (h *handler) CreateArticle(ctx context.Context, req *pb.Article, res *pb.ArticleResult) error {

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

func (h *handler) DeleteArticle(ctx context.Context, req *pb.Article, res *pb.ArticleResult) error {

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

func (h *handler) UpdateArticle(ctx context.Context, req *pb.Article, res *pb.ArticleResult) error {

	if req.Title == "" || req.Content == "" {
		return errors.BadRequest("", "Missing Param")
	}

	if req.Id == 0 {
		return errors.NotFound("", "Article Not Found")
	}

	err := h.usecase.UpdateArticle(ctx, req)
	if err != nil {
		return err
	}

	res.StatusCode = 200
	res.Message = "Successfully updated Article"

	return nil

}
