package service

import (
	"context"
	"errors"
	"strings"

	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryRedis"
)

type ICommentService interface {
	CreateComment(body string, userId, postId int) error
}

type CommentService struct {
	repositoryRedis *repositoryRedis.RepositoryRedis
}

func NewCommentService(repostiryRedis *repositoryRedis.RepositoryRedis) *CommentService {
	return &CommentService{
		repositoryRedis: repostiryRedis,
	}
}

func (s *CommentService) CreateComment(body string, userId, postId int) error {
	if strings.TrimSpace(body) == "" || userId == 0 || postId == 0 {
		return errors.New("body can't be empty")
	}
	ctx := context.Background()
	return s.repositoryRedis.PushComment(ctx, body, userId, postId)
}
