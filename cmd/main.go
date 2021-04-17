package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moll-y/ryoko/api/handler"
	"github.com/moll-y/ryoko/pkg/hero"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:toor@tcp(ryokodb:3306)/ryoko")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	heroRepository := hero.NewHeroMySQLRepository(db)
	heroService := hero.NewHeroService(heroRepository)
	heroHandler := handler.NewHeroHandler(heroService)

	r := mux.NewRouter()
	r.HandleFunc("/heroes/{ID:[0-9]+}", heroHandler.DeleteHero).Methods(http.MethodDelete)
	r.HandleFunc("/heroes/{ID:[0-9]+}", heroHandler.GetHero).Methods(http.MethodGet)
	r.HandleFunc("/heroes", heroHandler.GetHeroes).Methods(http.MethodGet)
	r.HandleFunc("/heroes", heroHandler.AddHero).Methods(http.MethodPost)
	r.HandleFunc("/heroes", heroHandler.UpdateHero).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":8080", r))
}
