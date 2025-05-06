package objects

import (
	"encoding/json"
	"errors"
)

// ParseData convierte los datos entrantes en un mapa de strings.
func ParseData(data []byte) (map[string]interface{}, error) {
	var parsedData map[string]interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return nil, errors.New("error al parsear datos")
	}
	return parsedData, nil
}
