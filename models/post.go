package models

import "time"

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

const (
	SAVED     Status = "saved"
	PUBLISHED Status = "published"
)

type Status string

type Post struct {
	ID         string
	Title      string
	Body       string
	Categories []string
	Tags       []string
	UserID     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     Status
	Views      int64
	Favourites int64
}
