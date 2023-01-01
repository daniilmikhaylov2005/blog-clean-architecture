package repositoryPostgres

import (
	"database/sql"

	"github.com/daniilmikhaylov2005/blog/internal/models"
)

type IPostRepository interface {
  InsertPost(post models.Post, userId int) (int, error)
}

type PostRepository struct {
  db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
  return &PostRepository{
    db: db,
  }
}

func (r *PostRepository) InsertPost(post models.Post, userId int) (int, error) {
  query := `INSERT INTO posts (title, body, user_id) VALUES ($1, $2, $3) RETURNING id`
  row := r.db.QueryRow(query, post.Title, post.Body, userId)
  
  var id int

  if err := row.Scan(&id); err != nil {
    return 0, err
  }

  return id, nil
}
