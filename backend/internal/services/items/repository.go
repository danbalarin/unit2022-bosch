package items

import (
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"gorm.io/gorm"
)

type itemsRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IItemsRepository {
	return &itemsRepository{
		db: db,
	}
}

func (repo *itemsRepository) insertItem(item *entity.Item) error {
	if err := repo.db.Create(item).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrItemAlreadyExists
			}
		}

		return errors.Wrap(err, "failed to insert item")
	}
	return nil
}

func (repo *itemsRepository) findItems() ([]entity.Item, error) {
	var items []entity.Item
	if err := repo.db.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find items")
	}
	return items, nil
}
