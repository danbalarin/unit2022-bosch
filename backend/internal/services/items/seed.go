package items

import (
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"log"
)

type IItemsDbSeeder interface {
	Seed() error
}

type itemsDbSeeder struct {
	svc IItemsService
}

func NewDbSeeder(svc IItemsService) IItemsDbSeeder {
	return &itemsDbSeeder{
		svc: svc,
	}
}

func (s *itemsDbSeeder) Seed() error {
	log.Println("Seeding items database")

	itemNames := []string{
		"Sekačka",
		"Pilka",
		"Šroubky s kulatou hlavou",
		"Šroubky s placatou hlavou",
		"Dřevěná laťka 2x10m",
		"Plíšek 10x10m",
	}

	for _, name := range itemNames {
		err := s.svc.createItem(&entity.Item{
			Name: name,
		})
		if err != nil && err != ErrItemAlreadyExists {
			return err
		}
	}

	log.Println("Seeding items database complete!")

	return nil
}
