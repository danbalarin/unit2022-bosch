package warehouse

import (
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"gorm.io/gorm"
)

type warehousesRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IWarehouseRepository {
	return &warehousesRepository{
		db: db,
	}
}

func (repo *warehousesRepository) insertWarehouse(warehouse *entity.Warehouse) error {
	if err := repo.db.Create(warehouse).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrWarehouseAlreadyExists
			}
		}

		return errors.Wrap(err, "failed to insert warehouse")
	}
	return nil
}

func (repo *warehousesRepository) findWarehouses() ([]entity.Warehouse, error) {
	var warehouses []entity.Warehouse
	if err := repo.db.Find(&warehouses).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find warehouses")
	}
	return warehouses, nil
}

func (repo *warehousesRepository) insertRoute(route *entity.Route) error {
	if err := repo.db.Create(route).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrRouteAlreadyExists
			}
		}

		return errors.Wrap(err, "failed to insert route")
	}
	return nil
}

func (repo *warehousesRepository) findRoutes() ([]*entity.Route, error) {
	var routes []*entity.Route
	if err := repo.db.Preload("Waypoints").Find(&routes).Error; err != nil {
		return nil, errors.Wrap(err, "failed to find routes")
	}
	return routes, nil
}

func (repo *warehousesRepository) insertWaypoint(waypoint *entity.Waypoint) error {
	if err := repo.db.Create(waypoint).Error; err != nil {
		return errors.Wrap(err, "failed to insert waypoint")
	}
	return nil
}
