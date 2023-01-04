package service

import (
	"errors"
	"strings"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
	"github.com/daniilmikhaylov2005/blog/pkg/utils"
  "time"
)

type IUserService interface {
  CreateUser(user models.User) error
  Signin(user models.User, accessToken string) (string, error)
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

func (s *UserService) Signin(user models.User, accessToken string) (string, error) {
  if strings.TrimSpace(user.Username) == "" || strings.TrimSpace(user.Password) == "" {
    return "", errors.New("Username or Password can't be empty")
  }

  userFromDb, err := s.repositoryPostgres.SelectUserByUsername(user.Username)
  if err != nil {
    return "", err
  }

  if !utils.CheckHashAndPassword(user.Password, userFromDb.Password) {
    return "", errors.New("Invalid password")
  }
  
  token, err := utils.CreateToken(userFromDb.ID, accessToken, time.Minute * 15)
  if err != nil {
    return "", err
  }

  return token, nil
}
