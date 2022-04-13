package entity

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);unique;not null"`
	Waypoints []*Waypoint
	Interval  int64 `json:"interval"` // in seconds
}
