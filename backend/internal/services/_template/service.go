package template

type ITemplateRepository interface {
}

type ITemplateService interface {
}

type templateService struct {
	repo ITemplateRepository
}

func NewService(repo ITemplateRepository) ITemplateService {
	return &templateService{
		repo: repo,
	}
}
