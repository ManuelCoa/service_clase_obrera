package ciudades

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectCiudad() (*sql.Rows, error) {
	query := "SELECT id_ciudad, id_estado, ciudad FROM config_geo_ciudades"
	return repository.FetchRows(query)
}

func selectCiudadID(id int) (*sql.Rows, error) {
	query := "SELECT id_ciudad, id_estado, ciudad FROM config_geo_ciudades WHERE id_ciudad = ?"
	return repository.FetchRows(query, id)
}

func insertCiudad(data Ciudad) (sql.Result, error) {
	query := "INSERT INTO config_geo_ciudades (id_estado, ciudad) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.EstadoID, data.CiudadNombre)
}

func updateCiudad(data Ciudad) (sql.Result, error) {
	query := "UPDATE config_geo_ciudades SET id_estado = ?, ciudad = ? WHERE id_ciudad = ?"
	return repository.ExecuteQuery(query, data.EstadoID, data.CiudadNombre, data.CiudadID)
}

func deleteCiudad(id int) (sql.Result, error) {
	query := "DELETE FROM config_geo_ciudades WHERE id_ciudad = ?"
	return repository.ExecuteQuery(query, id)
}
