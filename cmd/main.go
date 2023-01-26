package main

import (
	"log"

	"github.com/daniilmikhaylov2005/blog/internal/config"
	"github.com/daniilmikhaylov2005/blog/internal/handler"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryRedis"
	"github.com/daniilmikhaylov2005/blog/internal/service"
	"github.com/daniilmikhaylov2005/blog/server"
)

// @title Blog App API
// @version 1.0
// @description API Server for Blog website

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	config := config.InitConfig()

	redisDb := repositoryRedis.NewRepositoryRedis(config)
	postgresDb := repositoryPostgres.NewPostgres(config)
	service := service.NewService(postgresDb, redisDb)
	handler := handler.NewHandler(service, config)
	routes := handler.InitRoutes()

	srv := new(server.Server)
	if err := srv.Run(routes, config); err != nil {
		log.Printf("%s\n", err.Error())
	}
}
