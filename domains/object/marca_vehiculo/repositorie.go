package marca_vehiculo

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectMarcas() (*sql.Rows, error) {
	query := "SELECT id_marca_vehiculo, nombre_marca FROM config_marca_vehiculo"
	return repository.FetchRows(query)
}

func selectMarcasID(id int) (*sql.Rows, error) {
	query := "SELECT id_marca_vehiculo, nombre_marca FROM config_marca_vehiculo WHERE id_marca_vehiculo = ?"
	return repository.FetchRows(query, id)
}

func insertMarca(data MarcaVehiculo) (sql.Result, error) {
	query := "INSERT INTO config_marca_vehiculo (nombre_marca) VALUES (?)"
	return repository.ExecuteQuery(query, data.NombreMarca)
}

func updateMarca(data MarcaVehiculo) (sql.Result, error) {
	query := "UPDATE config_marca_vehiculo SET nombre_marca = ? WHERE id_marca_vehiculo = ?"
	return repository.ExecuteQuery(query, data.NombreMarca, data.MarcaVehiculoID)
}

func deleteMarca(id int) (sql.Result, error) {
	query := "DELETE FROM config_marca_vehiculo WHERE id_marca_vehiculo = ?"
	return repository.ExecuteQuery(query, id)
}
