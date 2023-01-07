package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("error while getting user id")
	}

	intUserId, ok := userId.(int)

	if !ok {
		return 0, errors.New("error while getting user id")
	}
	return intUserId, nil
}
