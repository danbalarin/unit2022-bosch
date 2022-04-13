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

func (repo *journeyRepository) insertItem(item *entity.Item) error {
	return nil
}
