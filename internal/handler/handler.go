package handler

import (
	_ "github.com/daniilmikhaylov2005/blog/docs"
	"github.com/daniilmikhaylov2005/blog/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Handler struct {
	services    *service.Service
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
			posts.POST("/", h.userIdentity, h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.GET("/:id", h.getPostById)
			posts.PUT("/:id", h.userIdentity, h.updatePostById)
			posts.DELETE("/:id", h.userIdentity, h.deletePostById)
			posts.POST("/comments/:id", h.userIdentity, h.createPostComment)
			posts.DELETE("/comments/:id", h.userIdentity, h.deletePostComment)
		}
		auth := api.Group("/auth")
		{
			auth.POST("/signup", h.signup)
			auth.POST("/signin", h.signin)
		}
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
