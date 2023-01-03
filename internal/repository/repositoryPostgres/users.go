package repositoryPostgres

import (
	"database/sql"

	"github.com/daniilmikhaylov2005/blog/internal/models"
)

type IUserRepository interface{
  InsertUser(user models.User) error
}

type UserRepository struct {
  db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
  return &UserRepository{
    db: db,
  }
}

func (r *UserRepository) InsertUser(user models.User) error {
  var id int

  tx, err := r.db.Begin()
  if err != nil {
    return err
  }

  query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id`
  row := tx.QueryRow(query, user.Username, user.Email, user.Password)

  if err := row.Scan(&id); err != nil {
    tx.Rollback()
    return err
  }

  tx.Commit()
  return nil
}
