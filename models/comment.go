package models

import "time"

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type Comment struct {
	ID         string
	UserID     string
	PostID     string
	Title      string
	Body       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Favourites int64
}
