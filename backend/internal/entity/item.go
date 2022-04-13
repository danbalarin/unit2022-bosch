package entity

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name string `json:"name" gorm:"type:varchar(255);unique;not null"`
}
