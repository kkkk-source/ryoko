package hero

type heroH2Repository struct {
	heroes []Hero
}

func NewHeroH2Repository() HeroRepository {
	return &heroRepository{
		[]Hero{
			{ID: 11, Name: "Dr Nice"},
			{ID: 12, Name: "Narco"},
			{ID: 13, Name: "Bombasto"},
			{ID: 14, Name: "Celeritas"},
			{ID: 15, Name: "Magneta"},
			{ID: 16, Name: "RubberMan"},
			{ID: 17, Name: "Dynama"},
			{ID: 18, Name: "Dr IQ"},
			{ID: 19, Name: "Magma"},
			{ID: 20, Name: "Tornado"},
		},
	}
}

func (r *heroH2Repository) Store(hero *Hero) (*Hero, error) {
	return nil, nil
}

func (r *heroH2Repository) Find(id int) (*Hero, error) {
	for i, h := range r.heroes {
		if h.ID == id {
			return h, nil
		}
	}
	return nil, error
}

func (r *heroH2Repository) FindAll() ([]*Hero, error) {
	return &r.heroes, nil
}

func (r *heroH2Repository) Upadte(hero *Hero) error {
	return nil
}

func (r *heroH2Repository) Destroy(id int) error {
	return nil
}