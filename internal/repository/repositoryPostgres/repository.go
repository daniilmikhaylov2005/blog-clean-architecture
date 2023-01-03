package repositoryPostgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type RepositoryPostgres struct {
	IPostRepository
  IUserRepository
}

func NewPostgres(config map[string]string) *RepositoryPostgres {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config["postgres_user"],
		config["postgres_password"],
		config["postgres_host"],
		config["postgres_port"],
		config["postgres_db"],
		config["postgres_ssl"])
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	return &RepositoryPostgres{
    IPostRepository: NewPostRepository(db),
    IUserRepository: NewUserRepository(db),
	}
}
