package config

import (
	"os"

	"log"

	"github.com/joho/godotenv"
)

func InitConfig() map[string]string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	config := make(map[string]string)
	config["postgres_user"] = os.Getenv("POSTGRES_USER")
	config["postgres_password"] = os.Getenv("POSTGRES_PASSWORD")
	config["postgres_host"] = os.Getenv("POSTGRES_HOST")
	config["postgres_port"] = os.Getenv("POSTGRES_PORT")
	config["postgres_db"] = os.Getenv("POSTGRES_DB")
	config["postgres_ssl"] = os.Getenv("POSTGRES_SSL")
  config["http_port"] = os.Getenv("HTTP_PORT")

	return config
}
