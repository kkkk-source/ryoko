package hero

type HeroService struct {
	repo HeroRepository
}

func NewHeroService(repo HeroRepository) *HeroService {
	return &HeroService{
		repo,
	}
}

func (s *HeroService) Save(hero *Hero) (*Hero, error) {
	heroCreated, err := s.repo.Store(hero)
	return heroCreated, err
}

func (s *HeroService) Find(id int64) (*Hero, error) {
	hero, err := s.repo.Find(id)
	return hero, err
}

func (s *HeroService) FindAll() ([]*Hero, error) {
	heroes, err := s.repo.FindAll()
	return heroes, err
}

func (s *HeroService) FindByName(name string) ([]*Hero, error) {
	heroes, err := s.repo.FindByName(name)
	return heroes, err
}

func (s *HeroService) Update(hero *Hero) error {
	return s.repo.Update(hero)
}

func (s *HeroService) Destroy(id int64) error {
	return s.repo.Destroy(id)
}
