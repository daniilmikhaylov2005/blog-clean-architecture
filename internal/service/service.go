package service

import (
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
)

type Service struct {
	IPostService
}

func NewService(postgres *repositoryPostgres.RepositoryPostgres) *Service {
  postService := NewPostService(postgres)
  return &Service{
    IPostService: postService,
  }
}
