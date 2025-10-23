package handler

import (
	"encoding/json"
	"net/http"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/repository"
)

// RandomPlaceNameHandler is a handler for random place name.
type RandomPlaceNameHandler struct {
	repo repository.PlaceNameRepository
}

// NewRandomPlaceNameHandler creates a new RandomPlaceNameHandler.
func NewRandomPlaceNameHandler(repo repository.PlaceNameRepository) *RandomPlaceNameHandler {
	return &RandomPlaceNameHandler{repo: repo}
}

// ServeHTTP handles the request.
func (h *RandomPlaceNameHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	placeName, err := h.repo.FindRandom()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(placeName); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
