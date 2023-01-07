package handler

import (
	"net/http"
	"strconv"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPostComment(c *gin.Context) {
	var postComment models.PostComment

	if err := c.BindJSON(&postComment); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	stringPostId := c.Param("id")
	postId, err := strconv.Atoi(stringPostId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	userId, err := getUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = h.services.CreateComment(postComment.Body, userId, postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, map[string]string{"Status": "Created"})
}
