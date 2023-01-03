package repositoryPostgres

import (
	"database/sql"

	"github.com/daniilmikhaylov2005/blog/internal/models"
)

type IPostRepository interface {
  InsertPost(post models.Post, userId int) (int, error)
  SelectAllPosts()([]models.Post, error)
  SelectPostById(id int) (models.Post, error)
  PutPost(post models.Post, postId int) (models.Post, error)
  DeletePost(postId, userId int) (int, error)
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
  tx, err := r.db.Begin()

  if err != nil {
    return 0, err
  }

  query := `INSERT INTO posts (title, body, user_id) VALUES ($1, $2, $3) RETURNING id`
  row := tx.QueryRow(query, post.Title, post.Body, userId)
  
  var id int

  if err := row.Scan(&id); err != nil {
    tx.Rollback()
    return 0, err
  }
  
  tx.Commit()
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

func (r *PostRepository) SelectPostById(id int) (models.Post, error) {
  var post models.Post

  query := `SELECT * FROM posts WHERE id=$1`
  row := r.db.QueryRow(query, id)

  if err := row.Scan(&post.ID, &post.Title, &post.Body, &post.UserId); err != nil {
    return models.Post{}, err
  }

  return post, nil
}

func (r *PostRepository) PutPost(post models.Post, postId int) (models.Post, error) {
  var id int

  tx, err := r.db.Begin()

  if err != nil {
    return models.Post{}, err
  }

  query := `UPDATE posts SET title=$1, body=$2, user_id=$3 WHERE id=$4 RETURNING id`
  row := tx.QueryRow(query, post.Title, post.Body, post.UserId, postId)
  
  if err := row.Scan(&id); err != nil {
    tx.Rollback()
    return models.Post{}, err
  }
  
  post.ID = id
  
  tx.Commit()

  return post, nil
}

func (r *PostRepository) DeletePost(postId, userId int) (int, error) {
  var deletedId int

  tx, err := r.db.Begin()

  if err != nil {
    return 0, err
  }

  query := `DELETE FROM posts WHERE id=$1 AND user_id=$2 RETURNING id`
  row := tx.QueryRow(query, postId, userId)

  if err := row.Scan(&deletedId); err != nil {
    tx.Rollback()

    return 0, err
  }

  tx.Commit()

  return deletedId, nil
}
