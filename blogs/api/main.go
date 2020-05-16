package main

import (
	"fmt"

	micro "github.com/micro/go-micro/v2"
	"github.com/whuangz/micro-blog/blogs-api/handler"

	pb "github.com/whuangz/micro-blog/blogs-api/proto"
)

func main() {
	srv := micro.NewService(
		micro.Name("micro-blogs-api-v1-blogs"),
		micro.Version("latest"),
	)

	srv.Init()

	h := handler.New(srv.Client())
	pb.RegisterBlogsHandler(srv.Server(), h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
