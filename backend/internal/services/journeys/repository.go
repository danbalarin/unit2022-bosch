package journeys

import "gorm.io/gorm"

type journeyRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IJourneyRepository {
	return &journeyRepository{
		db: db,
	}
}
