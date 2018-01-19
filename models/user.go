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
	Details    string
	UserType   UserType
	UserStatus UserStatus
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) Save() bool {
	return false
}

func (u *User) Update() bool {
	return false
}

func (u *User) Delete() bool {
	return false
}

func (u *User) FindAll() []User {
	return []User{}
}

func (u *User) FindByID() bool {
	return false
}

func (u *User) FindByUsername() bool {
	return false
}

func (u *User) ChangeStatus(userStatus UserStatus) bool {
	return false
}

func (u *User) ChangeType(userType UserType) bool {
	return false
}

func (u *User) ChangePassword(newPassword string) bool {
	return false
}
