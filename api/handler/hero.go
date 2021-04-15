package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/moll-y/ryoko/pkg/hero"
)

type heroHandler struct {
	service *hero.HeroService
}

func NewHeroHandler(service *hero.HeroService) *heroHandler {
	return &heroHandler{
		service,
	}
}

func (h *heroHandler) GetHero(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["ID"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	hero, err := h.service.Find(id)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Hero Not found")
		return
	}
	json.NewEncoder(w).Encode(hero)
}

func (h *heroHandler) GetHeroes(w http.ResponseWriter, r *http.Request) {
	heroes, _ := h.service.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}
