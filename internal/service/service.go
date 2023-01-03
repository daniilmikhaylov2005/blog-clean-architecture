package service

import (
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
)

type Service struct {
	IPostService
  IUserService
}

func NewService(postgres *repositoryPostgres.RepositoryPostgres) *Service {
  return &Service{
    IPostService: NewPostService(postgres),
    IUserService: NewUserService(postgres),
  }
}
