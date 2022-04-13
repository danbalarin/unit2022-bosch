package orm

import (
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"github.com/unit2022-bosch/teapot/backend/internal/services/auth"
	"gorm.io/gorm"
	"log"
)

type Migrator struct {
	db           *gorm.DB
	authDbSeeder auth.IAuthDbSeeder
}

func NewMigrator(db *gorm.DB, authDbSeeder auth.IAuthDbSeeder) *Migrator {
	return &Migrator{
		db:           db,
		authDbSeeder: authDbSeeder,
	}
}

var TABLES = []interface{}{
	&entity.User{},
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
	log.Println("Seeding database complete")

	return nil
}
