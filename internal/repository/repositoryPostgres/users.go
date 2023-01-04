package repositoryPostgres

import (
	"database/sql"

	"github.com/daniilmikhaylov2005/blog/internal/models"
)

type IUserRepository interface{
  InsertUser(user models.User) error
  SelectUserByUsername(username string) (models.User, error)
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

func (r *UserRepository) SelectUserByUsername(username string) (models.User, error) {
  var user models.User

  query := `SELECT * FROM users WHERE username=$1`
  row := r.db.QueryRow(query, username)

  if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
    return models.User{}, err
  }

  return user, nil
}
