package journeys

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"github.com/unit2022-bosch/teapot/backend/internal/services/warehouse"
	"log"
	"time"
)

type IJourneyWorker interface {
	Start() error
}

type journeyWorker struct {
	svcJourney   IJourneyService
	svcWarehouse warehouse.IWarehouseService
}

func NewWorker(svcJourney IJourneyService, svcWarehouse warehouse.IWarehouseService) IJourneyWorker {
	return &journeyWorker{
		svcJourney:   svcJourney,
		svcWarehouse: svcWarehouse,
	}
}

func (worker *journeyWorker) Start() error {
	go func() {
		ticker := time.NewTicker(time.Millisecond * 1000)
		for {
			select {
			case <-ticker.C:
				startTime := time.Now()
				err := worker.Tick()
				if err != nil {
					log.Println("WORKER ERROR:", err)
				}
				endTime := time.Now()
				log.Printf("Tick took %s", endTime.Sub(startTime))
			}
		}
	}()
	return nil
}

func (worker *journeyWorker) Tick() error {
	journeys, err := worker.svcJourney.GetAllJourneys()
	if err != nil {
		return err
	}

	routes, err := worker.svcWarehouse.GetRoutes()
	if err != nil {
		return err
	}

	for _, route := range routes {
		route.IsSpawned = false
	}

	for _, journey := range journeys {
		running, err := worker.processJourney(journey)
		if err != nil {
			return err
		}

		if !running {
			for _, route := range routes {
				if route.ID == journey.RouteID {
					route.IsSpawned = true
				}
			}
		}
	}

	for _, route := range routes {
		if route.IsSpawned {
			// Don't need to spawn again
			continue
		}

		_, err := worker.svcJourney.CreateJourney(route)
		if err != nil {
			return err
		}
	}

	return nil
}

func (worker *journeyWorker) processJourney(journey *entity.Journey) (running bool, err error) {
	now := time.Now()

	shouldRun := journey.DepartureTime.Before(now)

	if shouldRun {
		worker.svcJourney.DepartureJourney(journey)

		currentWaypoint := 0
		endTime := journey.DepartureTime

		// Check at which waypoint the journey currently is
		for i, waypoint := range journey.Route.Waypoints {
			waypointTime := endTime.Add(time.Second * time.Duration(waypoint.Duration))
			endTime = waypointTime

			if waypointTime.Before(now) {
				currentWaypoint = i + 1
			}
		}

		// END OF JOURNEY
		if currentWaypoint == len(journey.Route.Waypoints) {
			err := worker.svcJourney.FinishJourney(journey)
			if err != nil {
				return false, errors.WithStack(err)
			}
			return false, nil
		}

		// RUNNING JOURNEY
		err := worker.svcJourney.updateJourneyPlace(journey, currentWaypoint)
		if err != nil {
			return false, errors.WithStack(err)
		}
		return true, nil
	}

	return false, nil
}
