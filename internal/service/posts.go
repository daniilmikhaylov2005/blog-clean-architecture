package service

import (
	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
)

type IPostService interface {
  CreatePost(post models.Post, userId int) (int, error)
}

type PostService struct {
  repositoryPotgres *repositoryPostgres.RepositoryPostgres
}

func NewPostService(repositoryPostgres *repositoryPostgres.RepositoryPostgres) *PostService {
  return &PostService{
    repositoryPotgres: repositoryPostgres,
  }
}

func (s *PostService) CreatePost(post models.Post, userId int) (int, error) {
  return s.repositoryPotgres.InsertPost(post, userId)
}
