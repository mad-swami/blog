package repositories

import (
	"database/sql"
	
	"github.com/user/blog/models"
)

type AdminRepository interface {
	Create(admin *models.Admin) error
	GetByID(id int64) (*models.Admin, error)
	GetByUsername(username string) (*models.Admin, error)
	Update(admin *models.Admin) error
	Delete(id int64) error
}

type SQLiteAdminRepository struct {
	DB *sql.DB
}

func NewSQLiteAdminRepository(db *sql.DB) AdminRepository {
	return &SQLiteAdminRepository{DB: db}
}

func (r *SQLiteAdminRepository) Create(admin *models.Admin) error {
	query := `INSERT INTO admin (username, password_hash, display_name)
			  VALUES (?, ?, ?)`
	
	result, err := r.DB.Exec(query, admin.Username, admin.PasswordHash, admin.DisplayName)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	admin.ID = id
	return nil
}

func (r *SQLiteAdminRepository) GetByID(id int64) (*models.Admin, error) {
	query := `SELECT id, username, password_hash, display_name, created_at FROM admin WHERE id = ?`
	
	var admin models.Admin
	err := r.DB.QueryRow(query, id).Scan(
		&admin.ID, 
		&admin.Username, 
		&admin.PasswordHash, 
		&admin.DisplayName, 
		&admin.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &admin, nil
}

func (r *SQLiteAdminRepository) GetByUsername(username string) (*models.Admin, error) {
	query := `SELECT id, username, password_hash, display_name, created_at FROM admin WHERE username = ?`
	
	var admin models.Admin
	err := r.DB.QueryRow(query, username).Scan(
		&admin.ID, 
		&admin.Username, 
		&admin.PasswordHash, 
		&admin.DisplayName, 
		&admin.CreatedAt,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &admin, nil
}

func (r *SQLiteAdminRepository) Update(admin *models.Admin) error {
	query := `UPDATE admin SET username = ?, password_hash = ?, display_name = ? WHERE id = ?`
	
	_, err := r.DB.Exec(query, admin.Username, admin.PasswordHash, admin.DisplayName, admin.ID)
	return err
}

func (r *SQLiteAdminRepository) Delete(id int64) error {
	query := `DELETE FROM admin WHERE id = ?`
	
	_, err := r.DB.Exec(query, id)
	return err
}
