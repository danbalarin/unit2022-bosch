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
	//GetTimeOfJourneys(user *entity.User) (*time.Time, *time.Time, error)
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
	for _, journey := range journeys {
		if !journey.Departed {
			return journey, nil
		}
	}
	return nil, nil
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

/*func (j journeyService) GetTimeOfJourneys(user *entity.User) (time.Time, *time.Time, error) {
	warehouseId := user.WorkspaceID
	journeys, err := j.repo.findJourneysByWarehouse(warehouseId)
	if err != nil {
		return time.Now(), nil, errors.WithStack(err)
	}

	var journeyWaiting *entity.Journey = nil
	var journeyDeparted *entity.Journey = nil

	for _, journey := range journeys {
		if journey.Departed {
			journeyDeparted = journey
		} else {
			journeyWaiting = journey
		}
	}

	var awaitingTime *time.Time = nil
	if journeyDeparted != nil {
		awaitingTime = &journeyDeparted.DepartureTime
		for _, waypoint := range journeyDeparted.Route.Waypoints {
			lastTime := *awaitingTime
			awaitingTime = lastTime.Add(time.Second * time.Duration(waypoint.Duration))
			if waypoint.WarehouseID == warehouseId {
				break
			}
		}
	}

	return journeyWaiting.DepartureTime, awaitingTime, nil
}*/
