package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

// mockPlaceNameRepository is a mock implementation of PlaceNameRepository for testing.
type mockPlaceNameRepository struct {
	placeName model.PlaceName
	err       error
}

func (m *mockPlaceNameRepository) FindRandom() (model.PlaceName, error) {
	return m.placeName, m.err
}

func TestRandomPlaceNameHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		repo := &mockPlaceNameRepository{
			placeName: model.PlaceName{Name: "test", Yomi: "yomi"},
		}
		handler := NewRandomPlaceNameHandler(repo)

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
		handler := NewRandomPlaceNameHandler(repo)

		req := httptest.NewRequest(http.MethodGet, "/random", nil)
		rec := httptest.NewRecorder()

		handler.ServeHTTP(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d, but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
}
