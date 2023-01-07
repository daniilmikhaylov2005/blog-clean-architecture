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
