package service

import (
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryRedis"
)

type Service struct {
	IPostService
	IUserService
	ICommentService
}

func NewService(postgres *repositoryPostgres.RepositoryPostgres, redis *repositoryRedis.RepositoryRedis) *Service {
	return &Service{
		IPostService:    NewPostService(postgres),
		IUserService:    NewUserService(postgres),
		ICommentService: NewCommentService(redis),
	}
}
