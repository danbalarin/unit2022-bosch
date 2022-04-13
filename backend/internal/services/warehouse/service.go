package warehouse

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
)

var ErrWarehouseAlreadyExists = errors.New("warehouse already exists")
var ErrRouteAlreadyExists = errors.New("route already exists")

type IWarehouseRepository interface {
	insertWarehouse(warehouse *entity.Warehouse) error
	findWarehouses() ([]entity.Warehouse, error)

	insertRoute(route *entity.Route) error
	findRoutes() ([]*entity.Route, error)

	insertWaypoint(waypoint *entity.Waypoint) error
}

type IWarehouseService interface {
	createWarehouse(warehouse *entity.Warehouse) error
	GetWarehouses() ([]entity.Warehouse, error)

	createRoute(route *entity.Route) error
	GetRoutes() ([]*entity.Route, error)

	createWaypoint(route *entity.Waypoint) error
}

type warehouseService struct {
	repo IWarehouseRepository
}

func NewService(repo IWarehouseRepository) IWarehouseService {
	return &warehouseService{
		repo: repo,
	}
}

func (src *warehouseService) createWarehouse(warehouse *entity.Warehouse) error {
	return src.repo.insertWarehouse(warehouse)
}

func (src *warehouseService) GetWarehouses() ([]entity.Warehouse, error) {
	return src.repo.findWarehouses()
}

func (src *warehouseService) createRoute(route *entity.Route) error {
	return src.repo.insertRoute(route)
}

func (src *warehouseService) GetRoutes() ([]*entity.Route, error) {
	return src.repo.findRoutes()
}

func (src *warehouseService) createWaypoint(waypoint *entity.Waypoint) error {
	return src.repo.insertWaypoint(waypoint)
}
