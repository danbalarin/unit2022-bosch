package orm

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"github.com/unit2022-bosch/teapot/backend/internal/services/items"
	"github.com/unit2022-bosch/teapot/backend/internal/services/journeys"
	"github.com/unit2022-bosch/teapot/backend/internal/services/warehouse"
	"gorm.io/gorm"
	"log"
)

type Migrator struct {
	db                *gorm.DB
	authDbSeeder      auth.IAuthDbSeeder
	journeysDbSeeder  journeys.IJourneyDbSeeder
	itemsDbSeeder     items.IItemsDbSeeder
	warehouseDbSeeder warehouse.IWarehouseDbSeeder
}

func NewMigrator(
	db *gorm.DB,
	authDbSeeder auth.IAuthDbSeeder,
	journeysDbSeeder journeys.IJourneyDbSeeder,
	itemsDbSeeder items.IItemsDbSeeder,
	warehouseDbSeeder warehouse.IWarehouseDbSeeder,
) *Migrator {
	return &Migrator{
		db:                db,
		authDbSeeder:      authDbSeeder,
		journeysDbSeeder:  journeysDbSeeder,
		itemsDbSeeder:     itemsDbSeeder,
		warehouseDbSeeder: warehouseDbSeeder,
	}
}

var TABLES = []interface{}{
	&entity.User{},
	&entity.Item{},
	&entity.Route{},
	&entity.Journey{},
	&entity.RequestedItems{},
	&entity.Warehouse{},
	&entity.Waypoint{},
}

func (m *Migrator) Clean() error {
	err := m.db.Migrator().DropTable(TABLES...)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (m *Migrator) Seed() error {
	log.Println("Seeding database")

	err := m.Clean()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.db.AutoMigrate(TABLES...)
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.warehouseDbSeeder.Seed()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.authDbSeeder.Seed()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.itemsDbSeeder.Seed()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.journeysDbSeeder.Seed()
	if err != nil {
		return errors.WithStack(err)
	}

	log.Println("Seeding database complete")

	return nil
}
