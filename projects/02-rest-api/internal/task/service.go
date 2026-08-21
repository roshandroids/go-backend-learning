package task

// Service validates before delegating to the repository. It's thin here
// because Task has almost no business logic — the split still exists
// to keep HTTP concerns (handler.go) separate from storage concerns
// (repository.go).
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateRequest) (Task, error) {
	if err := validateCreateRequest(req); err != nil {
		return Task{}, err
	}
	return s.repo.Create(req.Title), nil
}

func (s *Service) Get(id string) (Task, bool) {
	return s.repo.Get(id)
}

func (s *Service) List(cursor string, limit int) Page {
	return s.repo.List(cursor, limit)
}
