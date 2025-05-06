package amonestaciones

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectAmonestaciones() (*sql.Rows, error) {
	query := "SELECT id_amonestacion, cedula_obrero, fecha_amonestacion, motivo, sancion FROM gestion_amonestaciones"
	return repository.FetchRows(query)
}

func selectAmonestacionID(id int) (*sql.Rows, error) {
	query := "SELECT id_amonestacion, cedula_obrero, fecha_amonestacion, motivo, sancion FROM gestion_amonestaciones WHERE id_amonestacion = ?"
	return repository.FetchRows(query, id)
}

func insertAmonestacion(data Amonestaciones) (sql.Result, error) {
	query := "INSERT INTO gestion_amonestaciones (cedula_obrero, fecha_amonestacion, motivo, sancion) VALUES (?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.FechaAmonestacion, data.Motivo, data.Sancion)
}

func updateAmonestacion(data Amonestaciones) (sql.Result, error) {
	query := "UPDATE gestion_amonestaciones SET cedula_obrero = ?, fecha_amonestacion = ?, motivo = ?, sancion = ? WHERE id_amonestacion = ?"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.FechaAmonestacion, data.Motivo, data.Sancion, data.AmonestacionesID)
}

func deleteAmonestacion(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_amonestaciones WHERE id_amonestacion = ?"
	return repository.ExecuteQuery(query, id)
}
