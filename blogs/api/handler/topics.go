package handler

import (
	"context"

	pb "github.com/whuangz/micro-blog/blogs-api/proto"
)

func (h *Handler) FetchTopics(ctx context.Context, req *pb.ListRequest, res *pb.ListResponse) error {

	return nil
}

func (h *Handler) CreateTopic(ctx context.Context, req *pb.Topic, res *pb.Response) error {

	return nil
}
