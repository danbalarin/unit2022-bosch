package items

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
)

var ErrItemAlreadyExists = errors.New("item already exists")

type IItemsRepository interface {
	insertItem(item *entity.Item) error
	findItems() ([]entity.Item, error)
}

type IItemsService interface {
	createItem(item *entity.Item) error
	GetItems() ([]entity.Item, error)
}

type itemsService struct {
	repo IItemsRepository
}

func NewService(repo IItemsRepository) IItemsService {
	return &itemsService{
		repo: repo,
	}
}

func (src *itemsService) createItem(item *entity.Item) error {
	return src.repo.insertItem(item)
}

func (src *itemsService) GetItems() ([]entity.Item, error) {
	return src.repo.findItems()
}
