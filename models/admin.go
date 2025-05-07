package models

import (
	"time"
)

type Admin struct {
	ID           int64
	Username     string
	PasswordHash string
	DisplayName  string
	CreatedAt    time.Time
}
