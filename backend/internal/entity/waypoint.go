package entity

import "gorm.io/gorm"

type Waypoint struct {
	gorm.Model

	Route     *Route
	RouteID   uint
	Storage   *Storage
	StorageID uint
	Duration  uint
}
