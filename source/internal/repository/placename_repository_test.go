package repository

import (
	"reflect"
	"testing"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

func TestInMemoryPlaceNameRepository_FindRandom(t *testing.T) {
	placeNames := []model.PlaceName{
		{ID: 1, Name: "test1", Yomi: "yomi1"},
		{ID: 2, Name: "test2", Yomi: "yomi2"},
	}
	repo := NewInMemoryPlaceNameRepository(placeNames)

	randomPlace, err := repo.FindRandom()
	if err != nil {
		t.Fatalf("FindRandom failed: %v", err)
	}

	found := false
	for _, p := range placeNames {
		if p.ID == randomPlace.ID && p.Name == randomPlace.Name && p.Yomi == randomPlace.Yomi {
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

func TestInMemoryPlaceNameRepository_FindAll(t *testing.T) {
	t.Run("returns all place names when repository is not empty", func(t *testing.T) {
		placeNames := []model.PlaceName{
			{ID: 1, Name: "test1", Yomi: "yomi1"},
			{ID: 2, Name: "test2", Yomi: "yomi2"},
		}
		repo := NewInMemoryPlaceNameRepository(placeNames)

		allPlaceNames, err := repo.FindAll()
		if err != nil {
			t.Fatalf("FindAll failed: %v", err)
		}

		if !reflect.DeepEqual(allPlaceNames, placeNames) {
			t.Errorf("Expected %+v, got %+v", placeNames, allPlaceNames)
		}
	})

	t.Run("returns an empty slice when repository is empty", func(t *testing.T) {
		repo := NewInMemoryPlaceNameRepository([]model.PlaceName{})
		allPlaceNames, err := repo.FindAll()
		if err != nil {
			t.Fatalf("FindAll failed: %v", err)
		}

		if len(allPlaceNames) != 0 {
			t.Errorf("Expected an empty slice, but got %+v", allPlaceNames)
		}
	})
}

func TestInMemoryPlaceNameRepository_FindByID(t *testing.T) {
	t.Run("returns place name when ID exists", func(t *testing.T) {
		placeNames := []model.PlaceName{
			{ID: 1, Name: "test1", Yomi: "yomi1"},
			{ID: 2, Name: "test2", Yomi: "yomi2"},
		}
		repo := NewInMemoryPlaceNameRepository(placeNames)

		placeName, err := repo.FindByID(1)
		if err != nil {
			t.Fatalf("FindByID failed: %v", err)
		}

		expected := placeNames[0]
		if !reflect.DeepEqual(placeName, expected) {
			t.Errorf("Expected %+v, got %+v", expected, placeName)
		}
	})

	t.Run("returns error when ID does not exist", func(t *testing.T) {
		placeNames := []model.PlaceName{
			{ID: 1, Name: "test1", Yomi: "yomi1"},
		}
		repo := NewInMemoryPlaceNameRepository(placeNames)

		_, err := repo.FindByID(999)
		if err == nil {
			t.Fatal("Expected an error for non-existent ID, but got nil")
		}
	})
}
