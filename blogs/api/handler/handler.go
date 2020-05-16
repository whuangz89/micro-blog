package handler

import (
	"github.com/micro/go-micro/v2/client"
	blogs "github.com/whuangz/micro-blog/blogs/proto"
)

// Handler is an object can process RPC requests
type handler struct {
	blogs blogs.BlogsService
}

// New returns an instance of Handler
func New(client client.Client) *handler {
	return &handler{
		blogs: blogs.NewBlogsService("micro-blog-v1-blogs:8080", client),
	}
}
