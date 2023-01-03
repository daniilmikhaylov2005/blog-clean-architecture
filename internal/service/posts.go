package service

import (
	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
  "errors"
)

type IPostService interface {
  CreatePost(post models.Post, userId int) (int, error)
  GetAllPosts() ([]models.Post, error)
  GetPostById(id int) (models.Post, error)
  UpdatePost(post models.Post, postId int) (models.Post, error)
  DeletePost(postId, userId int) (int, error)
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

func (s *PostService) UpdatePost(post models.Post, postId int) (models.Post, error) {
  if post.Body == "" || post.Title == "" || post.UserId == 0 {
    return models.Post{}, errors.New("Body, title or user id can't be empty.")
  }
  return s.repositoryPostgres.PutPost(post, postId)
}

func (s *PostService) DeletePost(postId, userId int) (int, error) {
  return s.repositoryPostgres.DeletePost(postId, userId)
}
