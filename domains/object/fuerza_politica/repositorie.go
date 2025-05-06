package fuerza_politica

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectFuerzas() (*sql.Rows, error) {
	query := "SELECT id_fuerza_politica, nombre_fuerza FROM config_fuerza_politica"
	return repository.FetchRows(query)
}

func selectFuerzasID(id int) (*sql.Rows, error) {
	query := "SELECT id_fuerza_politica, nombre_fuerza FROM config_fuerza_politica WHERE id_fuerza_politica = ?"
	return repository.FetchRows(query, id)
}

func insertFuerza(data FuerzaPolitica) (sql.Result, error) {
	query := "INSERT INTO config_fuerza_politica (nombre_fuerza) VALUES (?)"
	return repository.ExecuteQuery(query, data.NombreFuerza)
}

func updateFuerza(data FuerzaPolitica) (sql.Result, error) {
	query := "UPDATE config_fuerza_politica SET nombre_fuerza = ? WHERE id_fuerza_politica = ?"
	return repository.ExecuteQuery(query, data.NombreFuerza, data.FuerzaPoliticaID)
}

func deleteFuerza(id int) (sql.Result, error) {
	query := "DELETE FROM config_fuerza_politica WHERE id_fuerza_politica = ?"
	return repository.ExecuteQuery(query, id)
}
