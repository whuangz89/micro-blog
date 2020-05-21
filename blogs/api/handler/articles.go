package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
	pb "github.com/whuangz/micro-blog/blogs/proto"
)

func (h *Handler) FetchArticles(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	articles, err := h.repository.FetchArticles(req)
	if err != nil {
		return err
	}
	res.Articles = articles

	return nil
}

func (h *Handler) CreateArticle(ctx context.Context, req *pb.Article, res *pb.Response) error {

	if req.Title == "" || req.Content == "" || req.TopicId == 0 {
		return errors.BadRequest("", "Missing Param")
	}

	err := h.repository.CreateArticle(ctx, req)

	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteArticle(ctx context.Context, req *pb.Article, res *pb.Response) error {

	if req.Id == 0 {
		return errors.BadRequest("", "Missing Article ID")
	}

	if err := h.repository.DeleteArticle(ctx, req); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateArticle(ctx context.Context, req *pb.Article, res *pb.Response) error {

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

	return nil

}
