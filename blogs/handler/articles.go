package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
	pb "github.com/whuangz/micro-blog/blogs/proto"
	"github.com/whuangz/micro-blog/blogs/repository"
)

func (h *handler) FetchArticles(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	articles, err := h.repository.FetchArticles()
	if err != nil {
		return err
	}
	res.Articles = make([]*pb.Article, len(articles))
	for i, a := range articles {
		res.Articles[i] = h.serializeArticle(a)
	}

	return nil
}

func (h *handler) CreateArticle(ctx context.Context, req *pb.Article, res *pb.Response) error {

	if req.Title == "" || req.Content == "" {
		return errors.BadRequest("", "Missing Param")
	}

	err := h.repository.CreateArticle(ctx, req)

	if err != nil {
		return err
	}
	return nil
}

func (h *handler) DeleteArticle(ctx context.Context, req *pb.Article, res *pb.Response) error {

	if req.Id == 0 {
		return errors.BadRequest("", "Missing Article ID")
	}

	if err := h.repository.DeleteArticle(ctx, req); err != nil {
		return err
	}
	return nil
}

func (h *handler) UpdateArticle(ctx context.Context, req *pb.Article, res *pb.Response) error {

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

func (h *handler) serializeArticle(a *repository.Article) *pb.Article {
	return &pb.Article{
		Id:        a.ID,
		Title:     a.Title,
		Slug:      a.Slug,
		Content:   a.Content,
		TopicId:   a.TopicID,
		AuthorId:  a.AuthorID,
		Status:    a.Status,
		CreatedAt: a.CreatedAt.Unix(),
		UpdatedAt: a.UpdatedAt.Unix(),
	}
}
