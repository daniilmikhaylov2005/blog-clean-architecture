package service

import (
	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
)

type IPostService interface {
  CreatePost(post models.Post, userId int) (int, error)
  GetAllPosts() ([]models.Post, error)
  GetPostById(id int) (models.Post, error)
}

type PostService struct {
  repositoryPostgres *repositoryPostgres.RepositoryPostgres
}

func NewPostService(repositoryPostgres *repositoryPostgres.RepositoryPostgres) *PostService {
  return &PostService{
    repositoryPostgres: repositoryPostgres,
  }
}

func (s *PostService) CreatePost(post models.Post, userId int) (int, error) {
  return s.repositoryPostgres.InsertPost(post, userId)
}

func (s *PostService) GetAllPosts() ([]models.Post, error) {
  return s.repositoryPostgres.SelectAllPosts()
}

func (s *PostService) GetPostById(id int) (models.Post, error) {
  return s.repositoryPostgres.SelectPostById(id)
}
