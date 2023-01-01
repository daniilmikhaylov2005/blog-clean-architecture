package main

import (
	"github.com/daniilmikhaylov2005/blog/internal/config"
	"github.com/daniilmikhaylov2005/blog/internal/handler"
	"github.com/daniilmikhaylov2005/blog/internal/repository/repositoryPostgres"
	"github.com/daniilmikhaylov2005/blog/internal/service"
	"github.com/daniilmikhaylov2005/blog/server"
  "log"
)

func main() {
	config := config.InitConfig()

	postgresDb := repositoryPostgres.NewPostgres(config)
  service := service.NewService(postgresDb)
  handler := handler.NewHandler(service)
  routes := handler.InitRoutes()
  
  srv := new(server.Server)
  if err := srv.Run(routes, config); err != nil {
    log.Printf("%s\n", err.Error())
  }
}
