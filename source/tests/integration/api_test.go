package integration

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

func TestAPIGatewayIntegration(t *testing.T) {
	apiURL := os.Getenv("API_GATEWAY_URL")
	if apiURL == "" {
		t.Skip("API_GATEWAY_URL is not set, skipping integration test")
	}

	t.Run("random endpoint", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/random")
		if err != nil {
			t.Fatalf("Failed to send request to API Gateway: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var placeName model.PlaceName
		if err := json.NewDecoder(resp.Body).Decode(&placeName); err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if placeName.ID == 0 || placeName.Name == "" || placeName.Yomi == "" {
			t.Errorf("Expected non-empty id, name and yomi, but got %+v", placeName)
		}
	})

	t.Run("list endpoint", func(t *testing.T) {
		resp, err := http.Get(apiURL + "/list")
		if err != nil {
			t.Fatalf("Failed to send request to API Gateway: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
		}

		var placeNames []model.PlaceName
		if err := json.NewDecoder(resp.Body).Decode(&placeNames); err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if len(placeNames) == 0 {
			t.Error("Expected a non-empty list of place names, but got an empty list")
		}

		for _, p := range placeNames {
			if p.ID == 0 || p.Name == "" || p.Yomi == "" {
				t.Errorf("Expected non-empty id, name and yomi, but got %+v", p)
			}
		}
	})
}
