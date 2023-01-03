package service

import (
	"errors"
	"strings"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
	"github.com/daniilmikhaylov2005/blog/pkg/utils"
)

type IUserService interface {
  CreateUser(user models.User) error
}

type UserService struct {
  repositoryPostgres *repositoryPostgres.RepositoryPostgres
}

func NewUserService(repositoryPostgres *repositoryPostgres.RepositoryPostgres) *UserService {
  return &UserService{
    repositoryPostgres: repositoryPostgres,
  }
}

func (s *UserService) CreateUser(user models.User) error {
  if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
    return errors.New("Username, Email or Password can't be empty.")
  }

  hashedPassword, err := utils.CreateHash(user.Password)

  if err != nil {
    return err
  }

  user.Password = hashedPassword

  return s.repositoryPostgres.InsertUser(user)
}
