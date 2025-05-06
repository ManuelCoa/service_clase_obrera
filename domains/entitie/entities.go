package entitie

import (
	"database/sql"
	"encoding/json"
	"errors"
)

func ScanRows[T any](rows *sql.Rows, scanFunc func(*sql.Rows, *T) error) ([]T, error) {
	var data []T
	for rows.Next() {
		var dato T
		if err := scanFunc(rows, &dato); err != nil {
			return nil, err
		}
		data = append(data, dato)
	}
	return data, nil
}

func ParseJSONToStruct[T any](data interface{}) (T, error) {
	var dato T
	jsonData, err := json.Marshal(data)
	if err != nil {
		return dato, errors.New("error al convertir los datos a JSON")
	}

	if err := json.Unmarshal(jsonData, &dato); err != nil {
		return dato, errors.New("error al convertir los datos a la estructura")
	}

	return dato, nil
}
