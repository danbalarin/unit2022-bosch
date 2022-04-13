package journeys

import (
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"gorm.io/gorm"
)

type journeyRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IJourneyRepository {
	return &journeyRepository{
		db: db,
	}
}

func (repo *journeyRepository) insertItemToCart(item *entity.Item) error {
	return nil
}

func (repo *journeyRepository) findJourneys() ([]*entity.Journey, error) {
	var journeys []*entity.Journey
	res := repo.db.
		Preload("Route").
		Preload("Route.Waypoints").
		Preload("Route.Waypoints.Warehouse").
		Preload("ItemRequests")
	if err := res.Find(&journeys).Error; err != nil {
		return nil, err
	}
	return journeys, nil
}

func (repo *journeyRepository) insertJourney(journey *entity.Journey) error {
	return repo.db.Create(journey).Error
}

func (repo *journeyRepository) deleteJourney(id uint) error {
	return repo.db.Delete(&entity.Journey{}, id).Error
}

func (repo *journeyRepository) updateJourneyPlace(journeyID uint, place int) error {
	return repo.db.Model(&entity.Journey{}).Where("id = ?", journeyID).Update("place", place).Error
}
