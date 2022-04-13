package entity

import "gorm.io/gorm"

type Waypoint struct {
	gorm.Model

	Route       *Route
	RouteID     uint
	Warehouse   *Warehouse
	WarehouseID uint
	Duration    uint
}
