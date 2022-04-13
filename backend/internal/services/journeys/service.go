package journeys

import (
	"github.com/pkg/errors"
	"github.com/unit2022-bosch/teapot/backend/internal/entity"
	"log"
	"time"
)

type IJourneyRepository interface {
	findJourneys() ([]*entity.Journey, error)
	insertJourney(journey *entity.Journey) error
	deleteJourney(id uint) error
	updateJourneyPlace(journeyID uint, place int) error
}

type IJourneyService interface {
	GetAllJourneys() ([]*entity.Journey, error)
	CreateJourney(route *entity.Route) (*entity.Journey, error)
	FinishJourney(journey *entity.Journey) error
	updateJourneyPlace(journey *entity.Journey, place int) error
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
	err := j.repo.insertJourney(journey)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return journey, nil
}

func (j journeyService) FinishJourney(journey *entity.Journey) error {
	err := j.repo.deleteJourney(journey.ID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (j journeyService) updateJourneyPlace(journey *entity.Journey, place int) error {
	if journey.Place == place {
		return nil
	}
	log.Printf("update journey %d (%s) place to %d", journey.ID, journey.Route.Name, place)
	err := j.repo.updateJourneyPlace(journey.ID, place)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
