package repository

import (
	"testing"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

func TestInMemoryPlaceNameRepository_FindRandom(t *testing.T) {
	placeNames := []model.PlaceName{
		{Name: "test1", Yomi: "yomi1"},
		{Name: "test2", Yomi: "yomi2"},
	}
	repo := NewInMemoryPlaceNameRepository(placeNames)

	randomPlace, err := repo.FindRandom()
	if err != nil {
		t.Fatalf("FindRandom failed: %v", err)
	}

	found := false
	for _, p := range placeNames {
		if p.Name == randomPlace.Name && p.Yomi == randomPlace.Yomi {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Expected to find the random place in the original slice, but not found: %+v", randomPlace)
	}
}

func TestInMemoryPlaceNameRepository_FindRandom_Empty(t *testing.T) {
	repo := NewInMemoryPlaceNameRepository([]model.PlaceName{})
	_, err := repo.FindRandom()
	if err == nil {
		t.Fatal("Expected an error for empty place names, but got nil")
	}
}
