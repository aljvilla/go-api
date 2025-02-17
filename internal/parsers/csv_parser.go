package parsers

import (
	"encoding/csv"
	"errors"
	"io"
)

// ParseCSV procesa un archivo CSV y devuelve un array de filas como arrays de strings
func ParseCSV(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)
	var records [][]string

	// Leer todas las filas
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.New("error al leer el archivo CSV")
	}
	return records, nil
}
