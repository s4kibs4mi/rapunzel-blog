package models

import "time"

type Comment struct {
	ID              string
	UserID          string
	PostID          string
	ParentCommentID string
	Title           string
	Body            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Favourites      int64
}
