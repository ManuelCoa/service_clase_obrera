package solicitud_egreso

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectSolicitudEgreso() (*sql.Rows, error) {
	query := "SELECT solicitud_egreso_id, obrero_cedula, fecha_solicitud, motivo, estatus, observaciones FROM gestion_solicitud_egreso"
	return repository.FetchRows(query)
}

func selectSolicitudEgresoID(id int) (*sql.Rows, error) {
	query := "SELECT solicitud_egreso_id, obrero_cedula, fecha_solicitud, motivo, estatus, observaciones FROM gestion_solicitud_egreso WHERE solicitud_egreso_id = ?"
	return repository.FetchRows(query, id)
}

func insertSolicitudEgreso(data SolicitudEgreso) (sql.Result, error) {
	query := "INSERT INTO gestion_solicitud_egreso (obrero_cedula, fecha_solicitud, motivo, estatus, observaciones) VALUES (?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.FechaSolicitud, data.Motivo, data.Estatus, data.Observaciones)
}

func updateSolicitudEgreso(data SolicitudEgreso) (sql.Result, error) {
	query := "UPDATE gestion_solicitud_egreso SET obrero_cedula = ?, fecha_solicitud = ?, motivo = ?, estatus = ?, observaciones = ? WHERE solicitud_egreso_id = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.FechaSolicitud, data.Motivo, data.Estatus, data.Observaciones, data.SolicitudEgresoID)
}

func deleteSolicitudEgreso(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_solicitud_egreso WHERE solicitud_egreso_id = ?"
	return repository.ExecuteQuery(query, id)
}
