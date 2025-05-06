package vehiculo_obrero

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectVehiculoObrero() (*sql.Rows, error) {
	query := "SELECT vehiculo_id, marca_vehiculo_id, obrero_cedula, placa, modelo, color, nivel_gasolina FROM gestion_vehiculo_obrero"
	return repository.FetchRows(query)
}

func selectVehiculoObreroID(id int) (*sql.Rows, error) {
	query := "SELECT vehiculo_id, marca_vehiculo_id, obrero_cedula, placa, modelo, color, nivel_gasolina FROM gestion_vehiculo_obrero WHERE vehiculo_id = ?"
	return repository.FetchRows(query, id)
}

func insertVehiculoObrero(data VehiculoObrero) (sql.Result, error) {
	query := "INSERT INTO gestion_vehiculo_obrero (marca_vehiculo_id, obrero_cedula, placa, modelo, color, nivel_gasolina) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.MarcaVehiculoID, data.ObreroCedula, data.Placa, data.Modelo, data.Color, data.NivelGasolina)
}

func updateVehiculoObrero(data VehiculoObrero) (sql.Result, error) {
	query := "UPDATE gestion_vehiculo_obrero SET marca_vehiculo_id = ?, obrero_cedula = ?, placa = ?, modelo = ?, color = ?, nivel_gasolina = ? WHERE vehiculo_id = ?"
	return repository.ExecuteQuery(query, data.MarcaVehiculoID, data.ObreroCedula, data.Placa, data.Modelo, data.Color, data.NivelGasolina, data.VehiculoID)
}

func deleteVehiculoObrero(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_vehiculo_obrero WHERE vehiculo_id = ?"
	return repository.ExecuteQuery(query, id)
}
