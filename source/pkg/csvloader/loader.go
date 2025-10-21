package csvloader

import (
	"encoding/csv"
	"os"

	"github.com/t-ogawa/hokkaido-nandoku-api/internal/model"
)

// LoadPlaceNames loads place names from a CSV file.
func LoadPlaceNames(filePath string) ([]model.PlaceName, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2 // name, yomi
	
	// Skip header
	if _, err := reader.Read(); err != nil {
		return nil, err
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	placeNames := make([]model.PlaceName, 0, len(records))
	for _, record := range records {
		placeNames = append(placeNames, model.PlaceName{
			Name: record[0],
			Yomi: record[1],
		})
	}

	return placeNames, nil
}
