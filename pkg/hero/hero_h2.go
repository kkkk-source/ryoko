package hero

import "errors"

type heroH2Repository struct {
	heroes []*Hero
}

func NewHeroH2Repository() HeroRepository {
	return &heroH2Repository{
		[]*Hero{
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
			{ID: 21, Name: "Tornado"},
		},
	}
}

func (r *heroH2Repository) Store(hero *Hero) (*Hero, error) {
	// Find the bigger number and use it to create an unique one
	id := r.heroes[0].ID
	for _, h := range r.heroes {
		if id < h.ID {
			id = h.ID
		}
	}
	hero.ID = id + 1
	r.heroes = append(r.heroes, hero)
	return hero, nil
}

func (r *heroH2Repository) Find(id int64) (*Hero, error) {
	for _, h := range r.heroes {
		if h.ID == id {
			return h, nil
		}
	}
	return nil, errors.New("Hero Not found")
}

func (r *heroH2Repository) FindAll() ([]*Hero, error) {
	return r.heroes, nil
}

func (r *heroH2Repository) FindByName(name string) ([]*Hero, error) {
	var heroes []*Hero
	for _, h := range r.heroes {
		if name == h.Name {
			heroes = append(heroes, h)
		}
	}
	return heroes, nil
}

func (r *heroH2Repository) Update(hero *Hero) error {
	for _, h := range r.heroes {
		if h.ID == hero.ID {
			h.Name = hero.Name
			return nil
		}
	}
	return errors.New("Hero Not found")
}

func (r *heroH2Repository) Destroy(id int64) error {
	for i, h := range r.heroes {
		if h.ID == id {
			r.heroes[i] = r.heroes[len(r.heroes)-1]
			r.heroes = r.heroes[:len(r.heroes)-1]
			return nil
		}
	}
	return errors.New("Hero Not found")
}
