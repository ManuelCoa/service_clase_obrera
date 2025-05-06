package genero

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectGenero() (*sql.Rows, error) {
	query := "SELECT  id_genero, nombre_genero FROM config_genero"
	return repository.FetchRows(query)
}

func selectGeneroID(id int) (*sql.Rows, error) {
	query := "SELECT  id_genero, nombre_genero FROM config_genero WHERE id_genero = ?"
	return repository.FetchRows(query, id)
}

func insertGenero(data Genero) (sql.Result, error) {
	query := "INSERT INTO config_genero ( nombre_genero) VALUES (?)"
	return repository.ExecuteQuery(query,data.NombreGenero)
}

func updateGenero(data Genero) (sql.Result, error) {
	query := "UPDATE config_genero SET nombre_genero = ? WHERE id_genero = ?"
	return repository.ExecuteQuery(query,data.NombreGenero, data.GeneroID)
}

func deleteGenero(id int) (sql.Result, error) {
	query := "DELETE FROM config_genero WHERE id_genero = ?"
	return repository.ExecuteQuery(query, id)
}
