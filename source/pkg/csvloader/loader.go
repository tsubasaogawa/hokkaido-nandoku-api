package csvloader

import (
	"encoding/csv"
	"os"
	"strconv"

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
	reader.FieldsPerRecord = 3 // id, name, yomi

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
		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}
		placeNames = append(placeNames, model.PlaceName{
			ID:   id,
			Name: record[1],
			Yomi: record[2],
		})
	}

	return placeNames, nil
}
