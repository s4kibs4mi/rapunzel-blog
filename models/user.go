package models

import "time"

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
	ID         string
	Name       string
	Username   string
	Password   string
	details    string
	UserType   UserType
	UserStatus UserStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
