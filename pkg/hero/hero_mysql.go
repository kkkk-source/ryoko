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
	stmt, err := r.db.Prepare("INSERT INTO heroes(name) VALUES(?)")
	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(hero.Name)
	if err != nil {
		return nil, err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	hero.ID = lid
	return hero, nil
}

func (r heroMySQLRepository) Find(id int64) (*Hero, error) {
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
	stmt, err := r.db.Prepare("UPDATE heroes SET name=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(hero.Name, hero.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r heroMySQLRepository) Destroy(id int64) error {
	stmt, err := r.db.Prepare("DELETE FROM heroes WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
