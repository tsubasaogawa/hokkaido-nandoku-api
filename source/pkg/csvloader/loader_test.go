package csvloader

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

func TestLoadPlaceNames(t *testing.T) {
	// Create a temporary CSV file for testing
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	csvFile, err := os.Create(filepath.Join(tmpDir, "test.csv"))
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer csvFile.Close()

	csvData := `name,yomi
test1,yomi1
test2,yomi2
`
	if _, err := csvFile.WriteString(csvData); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	// Test loading the CSV file
	placeNames, err := LoadPlaceNames(csvFile.Name())
	if err != nil {
		t.Fatalf("LoadPlaceNames failed: %v", err)
	}

	expected := []model.PlaceName{
		{Name: "test1", Yomi: "yomi1"},
		{Name: "test2", Yomi: "yomi2"},
	}

	if len(placeNames) != len(expected) {
		t.Fatalf("Expected %d place names, but got %d", len(expected), len(placeNames))
	}

	for i, p := range placeNames {
		if p.Name != expected[i].Name || p.Yomi != expected[i].Yomi {
			t.Errorf("Expected %+v, but got %+v", expected[i], p)
		}
	}
}
