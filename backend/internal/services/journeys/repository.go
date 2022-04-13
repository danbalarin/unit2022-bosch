package journeys

import (
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"gorm.io/gorm"
)

type journeyRepository struct {
	db *gorm.DB
}

func (repo *journeyRepository) findJourneysByWarehouse(warehouseID uint) ([]*entity.Journey, error) {
	var journeys []*entity.Journey

	res := repo.db.
		Preload("Route").
		Joins("LEFT JOIN waypoints ON journeys.route_id = waypoints.route_id").
		Where("waypoints.warehouse_id = ?", warehouseID).
		Find(&journeys)
	if err := res.Error; err != nil {
		return nil, err
	}

	return journeys, nil
}

func NewRepository(db *gorm.DB) IJourneyRepository {
	return &journeyRepository{
		db: db,
	}
}

func (repo *journeyRepository) insertItemToJourney(requestedItems *entity.RequestedItems) error {
	return repo.db.Create(requestedItems).Error
}

func (repo *journeyRepository) findJourneys() ([]*entity.Journey, error) {
	var journeys []*entity.Journey
	res := repo.db.
		Preload("Route").
		Preload("Route.Waypoints").
		Preload("Route.Waypoints.Warehouse").
		Preload("ItemRequests").
		Preload("ItemRequests.Item")
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
	return repo.db.Model(&entity.Journey{}).
		Where("id = ?", journeyID).
		Update("place", place).
		Error
}

func (repo *journeyRepository) updateDepartureJourney(journeyID uint) error {
	return repo.db.Model(&entity.Journey{}).
		Where("id = ?", journeyID).
		Update("departed", "true").
		Error
}

func (repo *journeyRepository) setItemsArrived(journeyID uint, warehouseID uint) error {
	return repo.db.Model(&entity.RequestedItems{}).
		Where("journey_id = ?", journeyID).
		Where("warehouse_id = ?", warehouseID).
		Update("arrived", true).
		Error
}
