package repositoryPostgres

import (
	"database/sql"

	"github.com/daniilmikhaylov2005/blog/internal/models"
)

type IPostRepository interface {
  InsertPost(post models.Post, userId int) (int, error)
  SelectAllPosts()([]models.Post, error)
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

func (r *PostRepository) SelectAllPosts() ([]models.Post, error) {
  var posts []models.Post

  query := `SELECT * FROM posts`
  rows, err := r.db.Query(query)

  if err != nil {
    return []models.Post{}, err
  }
  
  defer rows.Close()

  for rows.Next() {
    var post models.Post
    
    if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.UserId); err != nil {
      return []models.Post{}, err
    }

    posts = append(posts, post)
  }

  return posts, nil
}
