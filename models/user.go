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
	Parent UserType = "parent"
	Family UserType = "family"
	Ghost  UserType = "ghost"
)

const (
	Registered UserStatus = "registered"
	Verified   UserStatus = "verified"
	Blocked    UserStatus = "blocked"
)

type UserType string
type UserStatus string

type User struct {
	ID         bson.ObjectId
	Name       string
	Username   string
	Email      string
	Password   string
	Details    string
	UserType   UserType
	UserStatus UserStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
