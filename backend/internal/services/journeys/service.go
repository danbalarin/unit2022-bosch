package journeys

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
