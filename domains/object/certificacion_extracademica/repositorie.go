package certificacion_extracademica

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectCertificaciones() (*sql.Rows, error) {
	query := "SELECT certificacion_id, numero_certificacion, fecha_egreso FROM gestion_certificacion_extracademica"
	return repository.FetchRows(query)
}

func selectCertificacionID(id int) (*sql.Rows, error) {
	query := "SELECT certificacion_id, numero_certificacion, fecha_egreso FROM gestion_certificacion_extracademica WHERE certificacion_id = ?"
	return repository.FetchRows(query, id)
}

func insertCertificacion(data CertificacionExtracademica) (sql.Result, error) {
	query := "INSERT INTO gestion_certificacion_extracademica (numero_certificacion, fecha_egreso) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NumeroCertificacion, data.FechaEgreso)
}

func updateCertificacion(data CertificacionExtracademica) (sql.Result, error) {
	query := "UPDATE gestion_certificacion_extracademica SET numero_certificacion = ?, fecha_egreso = ? WHERE certificacion_id = ?"
	return repository.ExecuteQuery(query, data.NumeroCertificacion, data.FechaEgreso, data.CertificacionID)
}

func deleteCertificacion(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_certificacion_extracademica WHERE certificacion_id = ?"
	return repository.ExecuteQuery(query, id)
}
