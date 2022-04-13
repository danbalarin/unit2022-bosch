package entity

import (
	"gorm.io/gorm"
	"time"
)

type Journey struct {
	gorm.Model
	Route         *Route
	RouteID       uint `gorm:"not null"`
	DepartureTime time.Time
	ItemRequests  []*RequestedItems `gorm:"foreignkey:JourneyID"`
	Departed      bool
	Place         int // 0 = at central warehouse, 1 = at first station, 2 = second, ...
}

type RequestedItems struct {
	gorm.Model
	Journey       *Journey
	JourneyID     uint `gorm:"not null"`
	Item          *Item
	ItemID        uint `gorm:"not null"`
	Counts        uint `json:"counts"`
	RequestedBy   *User
	RequestedByID uint `gorm:"not null"`
	Warehouse     *Warehouse
	WarehouseID   uint `gorm:"not null"`
	Arrived       bool
}
