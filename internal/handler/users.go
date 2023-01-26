package handler

import (
	"net/http"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary sign up
// @Tags auth
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body models.User true "user info"
// @Success 200 {object} models.SuccessResponse
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/auth/signup [post]
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

	c.JSON(http.StatusOK, models.SuccessResponse{Status: "User created"})
}

// @Summary sign in
// @Tags auth
// @Description auth user
// @ID auth-user
// @Accept json
// @Produce json
// @Param input body models.User true "user info"
// @Success 200 {object} models.TokenResponse
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/auth/signin [post]
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
