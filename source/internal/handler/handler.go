package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/repository"
)

// Handler is a handler for place names.
type Handler struct {
	repo repository.PlaceNameRepository
}

// NewHandler creates a new Handler.
func NewHandler(repo repository.PlaceNameRepository) *Handler {
	return &Handler{repo: repo}
}

// ServeHTTP handles the request.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if strings.HasSuffix(r.URL.Path, "/list") {
		placeNames, err := h.repo.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(placeNames); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if strings.HasSuffix(r.URL.Path, "/random") {
		placeName, err := h.repo.FindRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(placeName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if strings.HasPrefix(r.URL.Path, "/id/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/id/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		placeName, err := h.repo.FindByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(placeName); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.NotFound(w, r)
	}
}
