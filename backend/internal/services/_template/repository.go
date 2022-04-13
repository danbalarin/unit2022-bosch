package template

import "gorm.io/gorm"

type templateRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ITemplateRepository {
	return &templateRepository{
		db: db,
	}
}
