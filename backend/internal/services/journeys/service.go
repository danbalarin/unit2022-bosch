package journeys

import "github.com/pkg/errors"

var ErrRouteExists = errors.New("route already exists")
var ErrWaypointExists = errors.New("waypoint already exists")
var ErrRouteNotFound = errors.New("route not found")

type IJourneyRepository interface {
}

type IJourneyService interface {
}

type journeyService struct {
	repo IJourneyRepository
}

func NewService(repo IJourneyRepository) IJourneyService {
	return &journeyService{
		repo: repo,
	}
}
