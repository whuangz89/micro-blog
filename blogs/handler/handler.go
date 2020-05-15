package handler

import (
	"github.com/whuangz/micro-blog/blogs/repository"
)

type handler struct {
	repository repository.Service
}

func NewBlogHandler(repo repository.Service) *handler {
	return &handler{
		repository: repo,
	}
}
