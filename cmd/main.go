package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moll-y/ryoko/api/handler"
	"github.com/moll-y/ryoko/pkg/hero"
)

func main() {
	heroRepository := hero.NewHeroH2Repository()
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
