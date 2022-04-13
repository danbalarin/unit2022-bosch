package orm

import (
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"github.com/unit2022-bosch/teapot/backend/internal/services/items"
	"github.com/unit2022-bosch/teapot/backend/internal/services/journeys"
	"gorm.io/gorm"
	"log"
)

type Migrator struct {
	db               *gorm.DB
	authDbSeeder     auth.IAuthDbSeeder
	journeysDbSeeder journeys.IJourneyDbSeeder
	itemsDbSeeder    items.IItemsDbSeeder
}

func NewMigrator(
	db *gorm.DB,
	authDbSeeder auth.IAuthDbSeeder,
	journeysDbSeeder journeys.IJourneyDbSeeder,
	itemsDbSeeder items.IItemsDbSeeder,
) *Migrator {
	return &Migrator{
		db:               db,
		authDbSeeder:     authDbSeeder,
		journeysDbSeeder: journeysDbSeeder,
		itemsDbSeeder:    itemsDbSeeder,
	}
}

var TABLES = []interface{}{
	&entity.User{},
	&entity.Item{},
	&entity.Route{},
	&entity.Journey{},
	&entity.RequestedItems{},
	&entity.Storage{},
	&entity.Waypoint{},
}

func (m *Migrator) Clean() error {
	err := m.db.Migrator().DropTable(TABLES...)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migrator) Seed() error {
	log.Println("Seeding database")

	err := m.db.AutoMigrate(TABLES...)
	if err != nil {
		return err
	}

	err = m.authDbSeeder.Seed()
	if err != nil {
		return err
	}

	err = m.itemsDbSeeder.Seed()
	if err != nil {
		return err
	}

	err = m.journeysDbSeeder.Seed()
	if err != nil {
		return err
	}

	log.Println("Seeding database complete")

	return nil
}
