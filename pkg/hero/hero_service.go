package hero

type HeroService struct {
	repo HeroRepository
}

func NewHeroService(repo HeroRepository) *HeroService {
	return &HeroService{
		repo,
	}
}

func (s *HeroService) Find(id int) (*Hero, error) {
	hero, err := s.repo.Find(id)
	return hero, err
}

func (s *HeroService) FindAll() ([]*Hero, error) {
	return s.repo.FindAll()
}
