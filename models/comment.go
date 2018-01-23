package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

type Comment struct {
	ID         bson.ObjectId `bson:"_id"`
	UserID     bson.ObjectId `bson:"user_id"`
	PostID     bson.ObjectId `bson:"post_id"`
	Title      string        `bson:"title"`
	Body       string        `bson:"body"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at"`
	Favourites int64         `bson:"favourites"`
}
