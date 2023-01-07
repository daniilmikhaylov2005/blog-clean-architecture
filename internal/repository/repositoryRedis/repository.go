package repositoryRedis

import (
	"strconv"

	"github.com/go-redis/redis/v9"
)

type RepositoryRedis struct {
	ICommentsRepository
}

func NewRepositoryRedis(config map[string]string) *RepositoryRedis {
	redisDb, _ := strconv.Atoi(config["redis_db"])

	rdb := redis.NewClient(&redis.Options{
		Addr:     config["redis_addr"],
		Password: config["redis_password"],
		DB:       redisDb,
	})
	return &RepositoryRedis{
		ICommentsRepository: NewCommentRepository(rdb),
	}
}
