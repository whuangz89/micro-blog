package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
	pb "github.com/whuangz/micro-blog/blogs/proto"
	"github.com/whuangz/micro-blog/blogs/repository"
)

func (h *Handler) FetchTopics(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {
	topics, err := h.repository.FetchTopics(ctx, req)
	if err != nil {
		return err
	}

	res.Topics = make([]*pb.Topic, len(topics))
	for i, t := range topics {
		res.Topics[i] = h.serializeTopic(t)
	}
	return nil
}

func (h *Handler) CreateTopic(ctx context.Context, req *pb.Topic, res *pb.Response) error {

	if req.Title == "" || req.Color == "" {
		return errors.BadRequest("", "Missing Param")
	}

	err := h.repository.CreateTopic(ctx, req)

	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) GetTopic(ctx context.Context, req *pb.Topic, res *pb.Response) error {

	if req.Id == 0 {
		return errors.BadRequest("", "Missing Param")
	}

	err := h.repository.GetTopic(ctx, req)

	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) serializeTopic(t *repository.Topic) *pb.Topic {
	return &pb.Topic{
		Id:        t.ID,
		Title:     t.Title,
		Slug:      t.Slug,
		Tag:       t.Tag,
		Color:     t.Color,
		CreatedAt: t.CreatedAt.Unix(),
		UpdatedAt: t.UpdatedAt.Unix(),
	}
}
