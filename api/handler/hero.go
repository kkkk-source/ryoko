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

func (h *heroHandler) AddHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var hero hero.Hero
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	heroCreated, err := h.service.Save(&hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(heroCreated)
}

func (h *heroHandler) GetHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["ID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	hero, err := h.service.Find(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Hero Not found")
		return
	}
	json.NewEncoder(w).Encode(hero)
}

func (h *heroHandler) GetHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var heroes []*hero.Hero
	nameTerm := r.URL.Query().Get("name")
	if nameTerm != "" {
		heroes, _ = h.service.FindByName(nameTerm)
	} else {
		heroes, _ = h.service.FindAll()
	}
	json.NewEncoder(w).Encode(heroes)
}

func (h *heroHandler) UpdateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var hero hero.Hero
	err := json.NewDecoder(r.Body).Decode(&hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.service.Update(&hero)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h *heroHandler) DeleteHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["ID"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Bad Request")
		return
	}
	err = h.service.Destroy(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Hero Not found")
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
