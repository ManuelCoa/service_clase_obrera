package asistencia_biometrica

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectAsistenciasBiometricas() (*sql.Rows, error) {
	query := "SELECT id_asistencia_biometrica, cedula_obrero, identificacion_biometrica, fecha_registro, activa FROM gestion_asistencias_biometricas"
	return repository.FetchRows(query)
}

func selectAsistenciaBiometricaID(id int) (*sql.Rows, error) {
	query := "SELECT id_asistencia_biometrica, cedula_obrero, identificacion_biometrica, fecha_registro, activa FROM gestion_asistencias_biometricas WHERE id_asistencia_biometrica = ?"
	return repository.FetchRows(query, id)
}

func insertAsistenciaBiometrica(data AsistenciaBiometrica) (sql.Result, error) {
	query := "INSERT INTO gestion_asistencias_biometricas (cedula_obrero, identificacion_biometrica, fecha_registro, activa) VALUES (?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.IdentificacionBiometrica, data.FechaRegistro, data.Activa)
}

func updateAsistenciaBiometrica(data AsistenciaBiometrica) (sql.Result, error) {
	query := "UPDATE gestion_asistencias_biometricas SET cedula_obrero = ?, identificacion_biometrica = ?, fecha_registro = ?, activa = ? WHERE id_asistencia_biometrica = ?"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.IdentificacionBiometrica, data.FechaRegistro, data.Activa, data.AsistenciaBiometricaID)
}

func deleteAsistenciaBiometrica(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_asistencias_biometricas WHERE id_asistencia_biometrica = ?"
	return repository.ExecuteQuery(query, id)
}
