package repositories

import (
	"database/sql"
	
	"github.com/user/blog/models"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(id int64) (*models.Comment, error)
	GetByPostID(postID int64) ([]*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id int64) error
}

type SQLiteCommentRepository struct {
	DB *sql.DB
}

func NewSQLiteCommentRepository(db *sql.DB) CommentRepository {
	return &SQLiteCommentRepository{DB: db}
}

func (r *SQLiteCommentRepository) Create(comment *models.Comment) error {
	query := `INSERT INTO comments (post_id, commenter_name, content)
			  VALUES (?, ?, ?)`
	
	result, err := r.DB.Exec(query, comment.PostID, comment.CommenterName, comment.Content)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	comment.ID = id
	return nil
}

func (r *SQLiteCommentRepository) GetByID(id int64) (*models.Comment, error) {
	query := `SELECT id, post_id, commenter_name, content, created_at FROM comments WHERE id = ?`
	
	var comment models.Comment
	err := r.DB.QueryRow(query, id).Scan(
		&comment.ID,
		&comment.PostID,
		&comment.CommenterName,
		&comment.Content,
		&comment.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &comment, nil
}

func (r *SQLiteCommentRepository) GetByPostID(postID int64) ([]*models.Comment, error) {
	query := `SELECT id, post_id, commenter_name, content, created_at 
			  FROM comments 
			  WHERE post_id = ? 
			  ORDER BY created_at ASC`
	
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	comments := []*models.Comment{}
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.PostID,
			&comment.CommenterName,
			&comment.Content,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return comments, nil
}

func (r *SQLiteCommentRepository) Update(comment *models.Comment) error {
	query := `UPDATE comments 
			  SET commenter_name = ?, content = ? 
			  WHERE id = ?`
	
	_, err := r.DB.Exec(query, comment.CommenterName, comment.Content, comment.ID)
	return err
}

func (r *SQLiteCommentRepository) Delete(id int64) error {
	query := `DELETE FROM comments WHERE id = ?`
	
	_, err := r.DB.Exec(query, id)
	return err
}
