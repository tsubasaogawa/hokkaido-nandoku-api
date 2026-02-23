package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

// mockPlaceNameRepository is a mock implementation of PlaceNameRepository for testing.
type mockPlaceNameRepository struct {
	placeName  model.PlaceName
	placeNames []model.PlaceName
	err        error
}

func (m *mockPlaceNameRepository) FindRandom() (model.PlaceName, error) {
	return m.placeName, m.err
}

func (m *mockPlaceNameRepository) FindAll() ([]model.PlaceName, error) {
	return m.placeNames, m.err
}

func (m *mockPlaceNameRepository) FindByID(id int) (model.PlaceName, error) {
	if m.err != nil {
		return model.PlaceName{}, m.err
	}
	for _, p := range m.placeNames {
		if p.ID == id {
			return p, nil
		}
	}
	return model.PlaceName{}, errors.New("place name not found")
}

func TestRandomPlaceNameHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockPlaceNameRepository{
			placeName: model.PlaceName{ID: 1, Name: "test", Yomi: "yomi"},
		}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/random", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		var res model.PlaceName
		if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		if res.Name != "test" || res.Yomi != "yomi" {
			t.Errorf("expected %+v, but got %+v", repo.placeName, res)
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockPlaceNameRepository{
			err: errors.New("test error"),
		}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/random", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestPlaceNamesHandler_ListPlaceNames(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedPlaceNames := []model.PlaceName{
			{ID: 1, Name: "test1", Yomi: "yomi1"},
			{ID: 2, Name: "test2", Yomi: "yomi2"},
		}
		repo := &mockPlaceNameRepository{
			placeNames: expectedPlaceNames,
		}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/list", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		var res []model.PlaceName
		if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		if !reflect.DeepEqual(res, expectedPlaceNames) {
			t.Errorf("expected %+v, but got %+v", expectedPlaceNames, res)
		}
	})

	t.Run("error", func(t *testing.T) {
		repo := &mockPlaceNameRepository{
			err: errors.New("test error"),
		}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/list", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}

func TestHandler_GetPlaceNameByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectedPlaceNames := []model.PlaceName{
			{ID: 1, Name: "test1", Yomi: "yomi1"},
			{ID: 2, Name: "test2", Yomi: "yomi2"},
		}
		repo := &mockPlaceNameRepository{
			placeNames: expectedPlaceNames,
		}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/id/1", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d, but got %d", http.StatusOK, rec.Code)
		}

		var res model.PlaceName
		if err := json.NewDecoder(rec.Body).Decode(&res); err != nil {
			t.Fatalf("failed to decode response body: %v", err)
		}

		if !reflect.DeepEqual(res, expectedPlaceNames[0]) {
			t.Errorf("expected %+v, but got %+v", expectedPlaceNames[0], res)
		}
	})

	t.Run("not found", func(t *testing.T) {
		repo := &mockPlaceNameRepository{
			placeNames: []model.PlaceName{
				{ID: 1, Name: "test1", Yomi: "yomi1"},
			},
		}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/id/999", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("expected status code %d, but got %d", http.StatusNotFound, rec.Code)
		}
	})

	t.Run("invalid id", func(t *testing.T) {
		repo := &mockPlaceNameRepository{}
		handler := NewHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/id/invalid", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d, but got %d", http.StatusBadRequest, rec.Code)
		}
	})
}
