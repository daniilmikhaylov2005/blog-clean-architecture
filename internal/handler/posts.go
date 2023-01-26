package handler

import (
	"net/http"

	"strconv"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/gin-gonic/gin"
)

// @Summary Create
// @Tags posts
// @Description create post
// @Security ApiKeyAuth
// @ID create-post
// @Accept json
// @Produce json
// @Param input body models.Post true "post info"
// @Success 201 {integer} integer 1
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/posts/ [post]
func (h *Handler) createPost(c *gin.Context) {
	var post models.Post
	if err := c.BindJSON(&post); err != nil {
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

	postId, err := h.services.CreatePost(post, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]int{"post_id": postId})
}

// @Summary Get All
// @Tags posts
// @Description get all posts
// @ID get-post
// @Accept json
// @Produce json
// @Success 200 {array} models.Post
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/posts/ [get]
func (h *Handler) getAllPosts(c *gin.Context) {
	posts, err := h.services.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// @Summary Get By ID
// @Tags posts
// @Description get post by id
// @ID get-post-by-id
// @Accept json
// @Produce json
// @Success 200 {object} models.Post
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/posts/{id} [get]
func (h *Handler) getPostById(c *gin.Context) {
	paramId := c.Param("id")
	postId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	post, err := h.services.GetPostById(postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary Update
// @Tags posts
// @Description update post
// @Security ApiKeyAuth
// @ID update-post
// @Accept json
// @Produce json
// @Param input body models.Post true "post info"
// @Success 200 {object} models.Post
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/posts/{id} [put]
func (h *Handler) updatePostById(c *gin.Context) {
	stringId := c.Param("id")
	postId, err := strconv.Atoi(stringId)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	var post models.Post

	if err := c.BindJSON(&post); err != nil {
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

	post.UserId = userId

	updatedPost, err := h.services.UpdatePost(post, postId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedPost)
}

// @Summary Delete
// @Tags posts
// @Description delete post
// @Security ApiKeyAuth
// @ID delete-post
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 400,404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Failure default {object} models.ErrorResponse
// @Router /api/posts/{id} [delete]
func (h *Handler) deletePostById(c *gin.Context) {
	stringId := c.Param("id")
	postId, err := strconv.Atoi(stringId)

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

	deletedId, err := h.services.DeletePost(postId, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, map[string]int{"deleted_id": deletedId})
}
