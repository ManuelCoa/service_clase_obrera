package gestion_unoxdiez

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectGestionUnoxDiez() (*sql.Rows, error) {
	query := "SELECT unoxdiez_id, obrero_cedula, centro_votacion_id, nombres_apellidos, cedula, telefono, observacion FROM gestion_unoxdiez"
	return repository.FetchRows(query)
}

func selectGestionUnoxDiezID(id int) (*sql.Rows, error) {
	query := "SELECT unoxdiez_id, obrero_cedula, centro_votacion_id, nombres_apellidos, cedula, telefono, observacion FROM gestion_unoxdiez WHERE unoxdiez_id = ?"
	return repository.FetchRows(query, id)
}

func insertGestionUnoxDiez(data GestionUnoxDiez) (sql.Result, error) {
	query := "INSERT INTO gestion_unoxdiez (obrero_cedula, centro_votacion_id, nombres_apellidos, cedula, telefono, observacion) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.CentroVotacionID, data.NombresApellidos, data.Cedula, data.Telefono, data.Observacion)
}

func updateGestionUnoxDiez(data GestionUnoxDiez) (sql.Result, error) {
	query := "UPDATE gestion_unoxdiez SET obrero_cedula = ?, centro_votacion_id = ?, nombres_apellidos = ?, cedula = ?, telefono = ?, observacion = ? WHERE unoxdiez_id = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.CentroVotacionID, data.NombresApellidos, data.Cedula, data.Telefono, data.Observacion, data.UnoxDiezID)
}

func deleteGestionUnoxDiez(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_unoxdiez WHERE unoxdiez_id = ?"
	return repository.ExecuteQuery(query, id)
}
