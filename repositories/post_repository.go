package repositories

import (
	"database/sql"
	
	"github.com/user/blog/models"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetByID(id int64) (*models.Post, error)
	GetAll() ([]*models.Post, error)
	Update(post *models.Post) error
	Delete(id int64) error
}

type SQLitePostRepository struct {
	DB *sql.DB
}

func NewSQLitePostRepository(db *sql.DB) PostRepository {
	return &SQLitePostRepository{DB: db}
}

func (r *SQLitePostRepository) Create(post *models.Post) error {
	query := `INSERT INTO posts (title, content)
			  VALUES (?, ?)`
	
	result, err := r.DB.Exec(query, post.Title, post.Content)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	post.ID = id
	return nil
}

func (r *SQLitePostRepository) GetByID(id int64) (*models.Post, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM posts WHERE id = ?`
	
	var post models.Post
	err := r.DB.QueryRow(query, id).Scan(
		&post.ID, 
		&post.Title, 
		&post.Content, 
		&post.CreatedAt, 
		&post.UpdatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &post, nil
}

func (r *SQLitePostRepository) GetAll() ([]*models.Post, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM posts ORDER BY created_at DESC`
	
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	posts := []*models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.ID, 
			&post.Title, 
			&post.Content, 
			&post.CreatedAt, 
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return posts, nil
}

func (r *SQLitePostRepository) Update(post *models.Post) error {
	query := `UPDATE posts SET title = ?, content = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	
	_, err := r.DB.Exec(query, post.Title, post.Content, post.ID)
	return err
}

func (r *SQLitePostRepository) Delete(id int64) error {
	query := `DELETE FROM posts WHERE id = ?`
	
	_, err := r.DB.Exec(query, id)
	return err
}
