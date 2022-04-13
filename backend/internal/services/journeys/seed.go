package journeys

import (
	"log"
)

type IJourneyDbSeeder interface {
	Seed() error
}

type journeysDbSeeder struct {
	svc IJourneyService
}

func NewDbSeeder(svc IJourneyService) IJourneyDbSeeder {
	return &journeysDbSeeder{
		svc: svc,
	}
}

func (s *journeysDbSeeder) Seed() error {
	log.Println("Seeding journey database")
	log.Println("Seeding journey database complete!")

	return nil
}
