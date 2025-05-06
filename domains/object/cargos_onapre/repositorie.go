package cargos_onapre

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectCargoOnapre() (*sql.Rows, error) {
	query := "SELECT id_cagos_onapre, nombre_cargo, descripcion FROM config_cargos_onapre"
	return repository.FetchRows(query)
}

func selectCargoOnapreID(id int) (*sql.Rows, error) {
	query := "SELECT nombre_cargo, descripcion FROM config_geo_ciudades WHERE id_cargos_onapre = ?"
	return repository.FetchRows(query, id)
}

func insertCargoOnapre(data CargosOnapre) (sql.Result, error) {
	query := "INSERT INTO config_cargos_onapre (nombre_cargo, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.Nombre, data.Descripcion)
}

func updateCargoOnapre(data CargosOnapre) (sql.Result, error) {
	query := "UPDATE config_geo_ciudades SET id_estado = ?, ciudad = ? WHERE id_ciudad = ?"
	return repository.ExecuteQuery(query, data.Nombre, data.Descripcion, data.CargoOnapreID)
}

func deleteCargoOnapre(id int) (sql.Result, error) {
	query := "DELETE FROM config_cargos_onapre WHERE id_cargos_onapre = ?"
	return repository.ExecuteQuery(query, id)
}