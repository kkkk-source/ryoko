package hero

type HeroRepository interface {
	Store(hero *Hero) (*Hero, error)
	Find(id int64) (*Hero, error)
	FindAll() ([]*Hero, error)
	FindByName(name string) ([]*Hero, error)
	Update(hero *Hero) error
	Destroy(id int64) error
}
