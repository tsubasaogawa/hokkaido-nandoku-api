package repository

import (
	"errors"
	"math/rand"
	"time"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

// PlaceNameRepository is an interface for place name repository.
type PlaceNameRepository interface {
	FindRandom() (model.PlaceName, error)
	FindAll() ([]model.PlaceName, error)
}

// inMemoryPlaceNameRepository is an in-memory implementation of PlaceNameRepository.
type inMemoryPlaceNameRepository struct {
	placeNames []model.PlaceName
	rand       *rand.Rand
}

// NewInMemoryPlaceNameRepository creates a new inMemoryPlaceNameRepository.
func NewInMemoryPlaceNameRepository(placeNames []model.PlaceName) PlaceNameRepository {
	return &inMemoryPlaceNameRepository{
		placeNames: placeNames,
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// FindRandom returns a random place name.
func (r *inMemoryPlaceNameRepository) FindRandom() (model.PlaceName, error) {
	if len(r.placeNames) == 0 {
		return model.PlaceName{}, errors.New("no place names available")
	}
	return r.placeNames[r.rand.Intn(len(r.placeNames))], nil
}

// FindAll returns all place names.
func (r *inMemoryPlaceNameRepository) FindAll() ([]model.PlaceName, error) {
	return r.placeNames, nil
}
