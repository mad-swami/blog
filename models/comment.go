package models

import (
	"time"
)

type Comment struct {
	ID            int64
	PostID        int64
	CommenterName string
	Content       string
	CreatedAt     time.Time
}
