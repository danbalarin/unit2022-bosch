package entity

import "gorm.io/gorm"

type Warehouse struct {
	gorm.Model
	Name  string            `gorm:"type:varchar(100);not null;unique"`
	Items []*RequestedItems `gorm:"foreignkey:WarehouseID"`
}
