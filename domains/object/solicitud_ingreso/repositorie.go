package solicitud_ingreso

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectSolicitudIngreso() (*sql.Rows, error) {
	query := "SELECT solicitud_ingreso_id, institucion_id, cedula_solicitante, nombre_solicitante, apellido_solicitante, fecha_solicitud, estatus, observaciones FROM gestion_solicitud_ingreso"
	return repository.FetchRows(query)
}

func selectSolicitudIngresoID(id int) (*sql.Rows, error) {
	query := "SELECT solicitud_ingreso_id, institucion_id, cedula_solicitante, nombre_solicitante, apellido_solicitante, fecha_solicitud, estatus, observaciones FROM gestion_solicitud_ingreso WHERE solicitud_ingreso_id = ?"
	return repository.FetchRows(query, id)
}

func insertSolicitudIngreso(data SolicitudIngreso) (sql.Result, error) {
	query := "INSERT INTO gestion_solicitud_ingreso (institucion_id, cedula_solicitante, nombre_solicitante, apellido_solicitante, fecha_solicitud, estatus, observaciones) VALUES (?, ?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.InstitucionID, data.CedulaSolicitante, data.NombreSolicitante, data.ApellidoSolicitante, data.FechaSolicitud, data.Estatus, data.Observaciones)
}

func updateSolicitudIngreso(data SolicitudIngreso) (sql.Result, error) {
	query := "UPDATE gestion_solicitud_ingreso SET institucion_id = ?, cedula_solicitante = ?, nombre_solicitante = ?, apellido_solicitante = ?, fecha_solicitud = ?, estatus = ?, observaciones = ? WHERE solicitud_ingreso_id = ?"
	return repository.ExecuteQuery(query, data.InstitucionID, data.CedulaSolicitante, data.NombreSolicitante, data.ApellidoSolicitante, data.FechaSolicitud, data.Estatus, data.Observaciones, data.SolicitudIngresoID)
}

func deleteSolicitudIngreso(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_solicitud_ingreso WHERE solicitud_ingreso_id = ?"
	return repository.ExecuteQuery(query, id)
}
