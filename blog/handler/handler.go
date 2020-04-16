package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
	pb "github.com/whuangz/micro-blog/blog/proto"
	"github.com/whuangz/micro-blog/blog/repository"
)

type handler struct {
	repository repository.Service
}

func NewArticleHandler(repo repository.Service) *handler {
	return &handler{
		repository: repo,
	}
}

func (h *handler) FetchArticles(ctx context.Context, req *pb.ListArticleRequest, res *pb.ArticleResult) error {
	articles, err := h.repository.FetchArticles(ctx, req)
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

	err := h.repository.CreateArticle(ctx, req)

	if err != nil {
		return err
	}
	res.StatusCode = 200
	res.Message = "Success Inserted Article"
	return nil
}

func (h *handler) DeleteArticle(ctx context.Context, req *pb.Article, res *pb.ArticleResult) error {

	if req.Id == 0 {
		return errors.BadRequest("", "Missing Article ID")
	}

	if err := h.repository.DeleteArticle(ctx, req); err != nil {
		return err
	}
	res.StatusCode = 200
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

	err := h.repository.UpdateArticle(ctx, req)
	if err != nil {
		return err
	}

	res.StatusCode = 200
	res.Message = "Successfully updated Article"

	return nil

}
