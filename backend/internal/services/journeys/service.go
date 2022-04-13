package journeys

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"log"
	"time"
)

type AddItemsParams struct {
	ItemID      uint
	Count       uint
	User        *entity.User
	WarehouseID uint
}

type IJourneyRepository interface {
	findJourneys() ([]*entity.Journey, error)
	insertJourney(journey *entity.Journey) error
	deleteJourney(id uint) error
	updateJourneyPlace(journeyID uint, place int) error
	findJourneysByWarehouse(warehouseID uint) ([]*entity.Journey, error)
	insertItemToJourney(requestedItems *entity.RequestedItems) error
	updateDepartureJourney(journeyID uint) error
}

type IJourneyService interface {
	GetAllJourneys() ([]*entity.Journey, error)
	CreateJourney(route *entity.Route) (*entity.Journey, error)
	FinishJourney(journey *entity.Journey) error
	updateJourneyPlace(journey *entity.Journey, place int) error
	AddItemsToJourney(params AddItemsParams) error
	DepartureJourney(journey *entity.Journey) error
}

type journeyService struct {
	repo IJourneyRepository
}

func NewService(repo IJourneyRepository) IJourneyService {
	return &journeyService{
		repo: repo,
	}
}

func (j journeyService) GetAllJourneys() ([]*entity.Journey, error) {
	return j.repo.findJourneys()
}

func (j journeyService) CreateJourney(route *entity.Route) (*entity.Journey, error) {
	journey := &entity.Journey{
		RouteID:       route.ID,
		DepartureTime: time.Now().Add(time.Second * time.Duration(route.Interval)),
		ItemRequests:  nil,
		Place:         0,
	}

	log.Printf("[journeyService] CreateJourney: %v", journey)

	err := j.repo.insertJourney(journey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return journey, nil
}

func (j journeyService) DepartureJourney(journey *entity.Journey) error {
	if journey.Departed {
		return nil
	}

	err := j.repo.updateDepartureJourney(journey.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (j journeyService) FinishJourney(journey *entity.Journey) error {
	log.Printf("finishing journey %d (%s)", journey.ID, journey.Route.Name)
	err := j.repo.deleteJourney(journey.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (j journeyService) updateJourneyPlace(journey *entity.Journey, place int) error {
	if journey.Place == place && journey.Departed {
		return nil
	}
	log.Printf("update journey %d (%s) place to %d", journey.ID, journey.Route.Name, place)
	err := j.repo.updateJourneyPlace(journey.ID, place)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (svc journeyService) GetJourneyByWarehouse(warehouseID uint) (*entity.Journey, error) {
	journeys, err := svc.repo.findJourneysByWarehouse(warehouseID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(journeys) == 0 {
		return nil, nil
	}
	return journeys[0], nil
}

func (svc journeyService) AddItemsToJourney(params AddItemsParams) error {
	journey, err := svc.GetJourneyByWarehouse(params.WarehouseID)
	if err != nil {
		return errors.WithStack(err)
	}
	if journey == nil {
		return nil
	}

	log.Printf("add items to journey %d (%s): item_id=%d count=%d", journey.ID, journey.Route.Name, params.ItemID, params.Count)

	requestedItems := &entity.RequestedItems{
		JourneyID:     journey.ID,
		ItemID:        params.ItemID,
		RequestedByID: params.User.ID,
		Counts:        params.Count,
		WarehouseID:   params.WarehouseID,
	}

	err = svc.repo.insertItemToJourney(requestedItems)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
