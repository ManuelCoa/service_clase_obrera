package parroquias

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectParroquias() (*sql.Rows, error) {
	query := "SELECT id_parroquia, id_municipio, parroquia FROM config_geo_parroquias"
	return repository.FetchRows(query)
}

func selectParroquiasID(id int) (*sql.Rows, error) {
	query := "SELECT id_parroquia, id_municipio, parroquia FROM config_geo_parroquias WHERE id_parroquia = ?"
	return repository.FetchRows(query, id)
}

func insertParroquia(data Parroquia) (sql.Result, error) {
	query := "INSERT INTO config_geo_parroquias (id_municipio, parroquia) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.MunicipioID, data.ParroquiaNombre)
}

func updateParroquia(data Parroquia) (sql.Result, error) {
	query := "UPDATE config_geo_parroquias SET id_municipio = ?, parroquia = ? WHERE id_parroquia = ?"
	return repository.ExecuteQuery(query, data.MunicipioID, data.ParroquiaNombre, data.ParroquiaID)
}

func deleteParroquia(id int) (sql.Result, error) {
	query := "DELETE FROM config_geo_parroquias WHERE id_parroquia = ?"
	return repository.ExecuteQuery(query, id)
}
