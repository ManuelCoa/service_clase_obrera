package object

import (
	"encoding/json"
	"errors"
)

// ParseData convierte los datos entrantes en un mapa de strings.
func ParseData(data interface{}) (map[string]interface{}, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, errors.New("error al convertir datos a JSON")
	}
	var parsedData map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsedData); err != nil {
		return nil, errors.New("error al parsear datos")
	}
	return parsedData, nil
}
