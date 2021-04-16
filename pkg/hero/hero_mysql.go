package hero

import (
	"database/sql"
	"errors"
	"log"
)

type heroMySQLRepository struct {
	db *sql.DB
}

func NewHeroMySQLRepository(db *sql.DB) HeroRepository {
	return &heroMySQLRepository{
		db,
	}
}

func (r heroMySQLRepository) Store(hero *Hero) (*Hero, error) {
	return nil, nil
}

func (r heroMySQLRepository) Find(id int) (*Hero, error) {
	row, err := r.db.Query("SELECT * FROM heroes WHERE id = ?", id)
	defer row.Close()

	if err != nil {
		log.Fatal(err)
	}
	if !row.Next() {
		return nil, errors.New("Hero Not found")
	}

	var hero Hero
	err = row.Scan(&hero.ID, &hero.Name)
	if err != nil {
		log.Fatal(err)
	}
	return &hero, nil
}

func (r heroMySQLRepository) FindAll() ([]*Hero, error) {
	row, err := r.db.Query("SELECT * FROM heroes")
	defer row.Close()

	if err != nil {
		log.Fatal(err)
	}

	var heroes []*Hero
	for row.Next() {
		var hero Hero
		err = row.Scan(&hero.ID, &hero.Name)
		if err != nil {
			log.Fatal(err)
		}
		heroes = append(heroes, &hero)
	}
	return heroes, nil
}

func (r heroMySQLRepository) FindByName(name string) ([]*Hero, error) {
	row, err := r.db.Query("SELECT * FROM heroes")
	defer row.Close()

	if err != nil {
		log.Fatal(err)
	}

	var heroes []*Hero
	for row.Next() {
		var hero Hero
		err = row.Scan(&hero.ID, &hero.Name)
		if err != nil {
			log.Fatal(err)
		}
		if hero.Name == name {
			heroes = append(heroes, &hero)
		}
	}
	return heroes, nil
}

func (r heroMySQLRepository) Update(hero *Hero) error {
	return nil
}

func (r heroMySQLRepository) Destroy(id int) error {
	return nil
}
