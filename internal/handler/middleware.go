package handler

import (
	"strings"

	"net/http"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/daniilmikhaylov2005/blog/pkg/utils"
	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authHeader)

	if strings.TrimSpace(header) == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Empty Authorization header",
		})
    c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: "Wrong type of token",
		})
    c.Abort()
	  return
	}

	claims, err := utils.ParseToken(headerParts[1], h.accessToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
    c.Abort()
		return
	}

	c.Set(userCtx, claims.UserId)
}
