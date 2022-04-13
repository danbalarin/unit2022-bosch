package entity

import "gorm.io/gorm"

type Route struct {
	gorm.Model
	Name      string
	Waypoints []*Waypoint
	Interval  int64 `json:"interval"` // in seconds
}