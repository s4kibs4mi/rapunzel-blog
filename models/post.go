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

const (
	PostStatusSaved     Status = "saved"
	PostStatusPublished Status = "published"
)

type Status string

type Post struct {
	ID         bson.ObjectId `bson:"_id"`
	Title      string        `bson:"title"`
	Body       string        `bson:"body"`
	Categories []string      `bson:"categories"`
	Tags       []string      `bson:"tags"`
	UserID     bson.ObjectId `bson:"user_id"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at"`
	Status     Status        `bson:"status"`
	Views      int64         `bson:"views"`
	Favourites int64         `bson:"favourites"`
}
