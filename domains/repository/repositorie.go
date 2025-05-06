package repository

import (
	"database/sql"

	"claseobrera/persistence/interfaz"
)

func FetchRows(query string, args ...interface{}) (*sql.Rows, error) {
	db := interfaz.NewPersistence()
	rows, err := db.Read(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func ExecuteQuery(query string, args ...interface{}) (sql.Result, error) {
	db := interfaz.NewPersistence()
	result, err := db.Write(query, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
