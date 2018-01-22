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
	UserTypeParent UserType = "parent"
	UserTypeFamily UserType = "family"
	UserTypeGhost  UserType = "ghost"
)

const (
	UserStatusRegistered UserStatus = "registered"
	UserStatusVerified   UserStatus = "verified"
	UserStatusBlocked    UserStatus = "blocked"
)

type UserType string
type UserStatus string

type User struct {
	ID         bson.ObjectId `bson:"_id"`
	Name       string        `bson:"name"`
	Username   string        `bson:"username"`
	Email      string        `bson:"email"`
	Password   string        `bson:"password"`
	Details    string        `bson:"details"`
	UserType   UserType      `bson:"user_type"`
	UserStatus UserStatus    `bson:"user_status"`
	CreatedAt  time.Time     `bson:"created_at"`
	UpdatedAt  time.Time     `bson:"updated_at"`
}
