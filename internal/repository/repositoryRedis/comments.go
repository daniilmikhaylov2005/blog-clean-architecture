package repositoryRedis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type ICommentsRepository interface {
	PushComment(ctx context.Context, body string, userId, postId int) error
	DeleteComment(ctx context.Context, userId, postId int) error
}

type CommentRepository struct {
	rdb *redis.Client
}

func NewCommentRepository(rdb *redis.Client) *CommentRepository {
	return &CommentRepository{
		rdb: rdb,
	}
}

func (r *CommentRepository) PushComment(ctx context.Context, body string, userId, postId int) error {
	hashMap := fmt.Sprintf("comments_to_post_%d", postId)
	strUser := fmt.Sprintf("%d", userId)

	err := r.rdb.HSet(ctx, hashMap, strUser, body).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) DeleteComment(ctx context.Context, userId, postId int) error {
	hashMap := fmt.Sprintf("comments_to_post_%d", postId)
	strUser := fmt.Sprintf("%d", userId)

	err := r.rdb.HDel(ctx, hashMap, strUser).Err()
	if err != nil {
		return err
	}

	return nil
}
