package repositories

import (
	"database/sql"
	
	"github.com/user/blog/models"
)

type ImageRepository interface {
	Create(image *models.Image) error
	GetByID(id int64) (*models.Image, error)
	GetByPostID(postID int64) ([]*models.Image, error)
	Update(image *models.Image) error
	Delete(id int64) error
}

type SQLiteImageRepository struct {
	DB *sql.DB
}

func NewSQLiteImageRepository(db *sql.DB) ImageRepository {
	return &SQLiteImageRepository{DB: db}
}

func (r *SQLiteImageRepository) Create(image *models.Image) error {
	query := `INSERT INTO images (post_id, filename, file_path)
			  VALUES (?, ?, ?)`
	
	result, err := r.DB.Exec(query, image.PostID, image.Filename, image.FilePath)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	image.ID = id
	return nil
}

func (r *SQLiteImageRepository) GetByID(id int64) (*models.Image, error) {
	query := `SELECT id, post_id, filename, file_path, created_at FROM images WHERE id = ?`
	
	var image models.Image
	err := r.DB.QueryRow(query, id).Scan(
		&image.ID,
		&image.PostID,
		&image.Filename,
		&image.FilePath,
		&image.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &image, nil
}

func (r *SQLiteImageRepository) GetByPostID(postID int64) ([]*models.Image, error) {
	query := `SELECT id, post_id, filename, file_path, created_at 
			  FROM images 
			  WHERE post_id = ?`
	
	rows, err := r.DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	images := []*models.Image{}
	for rows.Next() {
		var image models.Image
		err := rows.Scan(
			&image.ID,
			&image.PostID,
			&image.Filename,
			&image.FilePath,
			&image.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		images = append(images, &image)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return images, nil
}

func (r *SQLiteImageRepository) Update(image *models.Image) error {
	query := `UPDATE images 
			  SET filename = ?, file_path = ? 
			  WHERE id = ?`
	
	_, err := r.DB.Exec(query, image.Filename, image.FilePath, image.ID)
	return err
}

func (r *SQLiteImageRepository) Delete(id int64) error {
	query := `DELETE FROM images WHERE id = ?`
	
	_, err := r.DB.Exec(query, id)
	return err
}
