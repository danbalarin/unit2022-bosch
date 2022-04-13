package warehouse

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"log"
)

type IWarehouseDbSeeder interface {
	Seed() error
}

type warehouseDbSeeder struct {
	svc IWarehouseService
}

func NewDbSeeder(svc IWarehouseService) IWarehouseDbSeeder {
	return &warehouseDbSeeder{
		svc: svc,
	}
}

func (s *warehouseDbSeeder) Seed() error {
	log.Println("Seeding warehouse database")
	warehouses := []*entity.Warehouse{
		{Name: "Apple Warehouse"},
		{Name: "Banana Warehouse"},
		{Name: "Cherry Warehouse"},
		{Name: "Date Warehouse"},
		{Name: "Elderberry Warehouse"},
		{Name: "Fig Warehouse"},
		{Name: "Grape Warehouse"},
	}

	for _, warehouse := range warehouses {
		s.svc.createWarehouse(warehouse)
		/*if err != nil || err != ErrWarehouseAlreadyExists {
			return errors.WithStack(err)
		}*/
	}

	log.Println("Seeding routes database")

	routeGolden := &entity.Route{
		Name:     "Zlatá cesta",
		Interval: 300,
	}
	err := s.svc.createRoute(routeGolden)
	if err != nil {
		if err != ErrRouteAlreadyExists {
			return errors.WithStack(err)
		}
	} else {
		s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeGolden.ID,
			WarehouseID: warehouses[0].ID,
			Duration:    40,
		})
		s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeGolden.ID,
			WarehouseID: warehouses[1].ID,
			Duration:    60,
		})
		s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeGolden.ID,
			WarehouseID: warehouses[2].ID,
			Duration:    40,
		})
	}

	routeSilver := &entity.Route{
		Name:     "Sametová cesta",
		Interval: 30,
	}
	err = s.svc.createRoute(routeSilver)
	if err != nil {
		if err != ErrRouteAlreadyExists {
			return errors.WithStack(err)
		}
	} else {
		log.Println(&entity.Waypoint{
			RouteID:     routeSilver.ID,
			WarehouseID: warehouses[3].ID,
			Duration:    2,
		})
		err := s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeSilver.ID,
			WarehouseID: warehouses[3].ID,
			Duration:    2,
		})
		if err != nil {
			return errors.WithStack(err)
		}
		s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeSilver.ID,
			WarehouseID: warehouses[4].ID,
			Duration:    4,
		})
		s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeSilver.ID,
			WarehouseID: warehouses[5].ID,
			Duration:    6,
		})
		s.svc.createWaypoint(&entity.Waypoint{
			RouteID:     routeSilver.ID,
			WarehouseID: warehouses[6].ID,
			Duration:    8,
		})
	}

	log.Println("Seeding warehouse database complete!")

	return nil
}
