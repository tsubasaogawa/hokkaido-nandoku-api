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

		if placeName.Name == "" || placeName.Yomi == "" {
			t.Errorf("Expected non-empty name and yomi, but got %+v", placeName)
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

		var body struct {
			PlaceNames []model.PlaceName `json:"placenames"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("Failed to decode response body: %v", err)
		}

		if len(body.PlaceNames) == 0 {
			t.Error("Expected a non-empty list of place names, but got an empty list")
		}

		for _, p := range body.PlaceNames {
			if p.Name == "" || p.Yomi == "" {
				t.Errorf("Expected non-empty name and yomi, but got %+v", p)
			}
		}
	})
}
