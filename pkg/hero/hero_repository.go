package hero

type HeroRepository interface {
	Store(hero *Hero) (*Hero, error)
	Find(id int) (*Hero, error)
	FindAll() ([]*Hero, error)
	FindByName(name string) ([]*Hero, error)
	Update(hero *Hero) error
	Destroy(id int) error
}
