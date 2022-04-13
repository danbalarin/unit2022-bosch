package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email             string `gorm:"type:varchar(200);unique"`
	PlaintextPassword string `gorm:"-"` // only used to create new user, then newer used
	Password          string // here is stored real password (in hash)
	Role              UserRole

	Workspace   *Storage
	WorkspaceID uint
}

type UserRole int

const (
	Role_User  UserRole = 0
	Role_Admin UserRole = 1
)

func (role *UserRole) IsUser() bool {
	return role != nil
}

func (role *UserRole) IsAdmin() bool {
	return *role == Role_Admin
}
