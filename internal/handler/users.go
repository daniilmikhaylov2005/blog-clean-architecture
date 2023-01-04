package handler

import (
	"net/http"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signup(c *gin.Context) {
  var user models.User
  
  if err := c.BindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  err := h.services.CreateUser(user)

  if err != nil {
    c.JSON(http.StatusInternalServerError, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, map[string]string{"Status": "User created"})
}

func (h *Handler) signin(c *gin.Context) {
  var user models.User

  if err := c.BindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  token, err := h.services.Signin(user, h.accessToken)

  if err != nil {
    c.JSON(http.StatusInternalServerError, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, map[string]string{"token": token})
}
