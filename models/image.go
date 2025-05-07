package models

import (
	"time"
)

type Image struct {
	ID        int64
	PostID    int64
	Filename  string
	FilePath  string
	CreatedAt time.Time
}
