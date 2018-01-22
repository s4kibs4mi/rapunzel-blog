package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	ID           bson.ObjectId `bson:"_id"`
	UserID       bson.ObjectId `bson:"user_id"`
	AccessToken  string        `bson:"access_token"`
	RefreshToken string        `bson:"refresh_token"`
	CreatedAt    time.Time     `bson:"created_at"`
	ExpiredAt    time.Time     `bson:"updated_at"`
}
