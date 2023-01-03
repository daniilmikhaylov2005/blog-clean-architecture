package handler

import (
	"github.com/daniilmikhaylov2005/blog/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
  accessToken string
}

func NewHandler(services *service.Service, config map[string]string) *Handler {
	return &Handler{services: services, accessToken: config["access_token"]}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.GET("/:id", h.getPostById)
			posts.PUT("/:id", h.updatePostById)
			posts.DELETE("/:id", h.deletePostById)
		}
    auth := api.Group("/auth")
    {
      auth.POST("/signup", h.signup)
      auth.POST("/signin", h.signin)
    }
	}

	return router
}
