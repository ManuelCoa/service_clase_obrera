package cargo_publico

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectCargo() (*sql.Rows, error) {
	query := "SELECT id_cargo, nombre, descripcion configuracion FROM config_cargo_publico"
	return repository.FetchRows(query)
}

func selectCargoID(id int) (*sql.Rows, error) {
	query := "SELECT id_cargo, nombre, descripcion FROM config_cargo_publico WHERE id_cargo = ?"
	return repository.FetchRows(query, id)
}

func insertCargo(data CargoPublico) (sql.Result, error) {
	query := "INSERT INTO config_cargo_publico (nombre, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query,data.Nombre, data.Nombre, data.Descripcion)
}

func updateCargo(data CargoPublico) (sql.Result, error) {
	query := "UPDATE config_cargo_publico SET nombre = ?, descripcion = ? WHERE id_cargo = ?"
	return repository.ExecuteQuery(query, data.Nombre, data.Descripcion, data.CargoID)
}

func deleteCargo(id int) (sql.Result, error) {
	query := "DELETE FROM config_cargo_publico WHERE id_cargo = ?"
	return repository.ExecuteQuery(query, id)
}
