package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	ID           bson.ObjectId
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
	ExpiredAt    time.Time
}
