package handler

import (
	"github.com/whuangz/micro-blog/blogs/repository"
)

type Handler struct {
	repository repository.Service
}

func NewBlogHandler(repo repository.Service) *Handler {
	return &Handler{
		repository: repo,
	}
}
