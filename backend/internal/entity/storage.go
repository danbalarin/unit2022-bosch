package entity

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	Name    string
	Workers []*User `gorm:"foreignKey:WorkspaceID"`
}
