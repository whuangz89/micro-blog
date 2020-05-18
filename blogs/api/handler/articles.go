package handler

import (
	"context"
	"time"

	"github.com/micro/go-micro/v2/errors"

	pb "github.com/whuangz/micro-blog/blogs-api/proto"
	blogs "github.com/whuangz/micro-blog/blogs/proto"
)

func (h *Handler) Fetch(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	rawArticlesRes, err := h.blogs.FetchArticles(ctx, &blogs.ListRequest{})
	if err != nil {
		return err
	}
	res.Articles = make([]*pb.Article, len(rawArticlesRes.Articles))
	for i, a := range rawArticlesRes.Articles {
		res.Articles[i] = h.serializeArticle(a)
	}

	return nil
}

func (h *Handler) Create(ctx context.Context, req *pb.Article, res *pb.Response) error {

	if req.Title == "" || req.Content == "" {
		return errors.BadRequest("", "Missing Param")
	}

	// _, err := h.blogs.CreateArticle(ctx, req)

	// if err != nil {
	// 	return err
	// }

	res.Error.Code = 200
	return nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.Article, res *pb.Response) error {

	return nil
}

func (h *Handler) Update(ctx context.Context, req *pb.Article, res *pb.Response) error {

	return nil

}

func (h *Handler) getTopicForArticle(a *blogs.Article) (*pb.Topic, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	rsp, err := h.blogs.GetTopic(ctx, &blogs.Topic{Id: a.TopicId})
	if err != nil {
		return nil, err
	}

	return &pb.Topic{
		Id:    rsp.Topic.Id,
		Title: rsp.Topic.Title,
		Slug:  rsp.Topic.Slug,
	}, nil

}

func (h *Handler) serializeArticle(a *blogs.Article) *pb.Article {

	res := &pb.Article{
		Id:        a.Id,
		Title:     a.Title,
		Slug:      a.Slug,
		Content:   a.Content,
		AuthorId:  a.AuthorId,
		Status:    a.Status,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}

	if topic, err := h.getTopicForArticle(a); err == nil {
		res.Topic = topic
	}

	return res

}
