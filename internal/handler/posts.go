package handler

import (
	"net/http"

	"strconv"

	"github.com/daniilmikhaylov2005/blog/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createPost(c *gin.Context) {
  var post models.Post
  if err := c.BindJSON(&post); err != nil {
    c.JSON(http.StatusBadRequest, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  fakeUserId := 1

  postId, err := h.services.CreatePost(post, fakeUserId)

  if err != nil {
    c.JSON(http.StatusInternalServerError, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  c.JSON(http.StatusOK, map[string]int{"post_id": postId})
}

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

func (h *Handler) getPostById(c *gin.Context) {
  paramId:= c.Param("id")
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

  fakeUserId := 1

  post.UserId = fakeUserId
  
  updatedPost, err := h.services.UpdatePost(post, postId)

  if err != nil {
    c.JSON(http.StatusInternalServerError, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

	c.JSON(http.StatusOK, updatedPost)
}

func (h *Handler) deletePostById(c *gin.Context) {
  stringId := c.Param("id")
  postId, err := strconv.Atoi(stringId)

  if err != nil {
    c.JSON(http.StatusBadRequest, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

  fakeUserId := 1

  deletedId, err := h.services.DeletePost(postId, fakeUserId)

  if err != nil {
    c.JSON(http.StatusInternalServerError, models.ErrorResponse{
      Error: err.Error(),
    })
    return
  }

	c.JSON(http.StatusOK, map[string]int{"deleted_id": deletedId})
}
