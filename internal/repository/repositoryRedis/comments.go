package repositoryRedis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

type ICommentsRepository interface {
	PushComment(ctx context.Context, body string, userId, postId int) error
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
	strPostId := fmt.Sprintf("comments_of_post_%d", postId)
	comment := fmt.Sprintf("%d:%s", userId, body)

	err := r.rdb.LPush(ctx, strPostId, comment).Err()
	if err != nil {
		return err
	}

	return nil
}
