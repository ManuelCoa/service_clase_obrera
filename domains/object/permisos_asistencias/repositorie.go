package permisos_asistencias

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectPermisoAsistencia() (*sql.Rows, error) {
	query := "SELECT permiso_asistencia_id, obrero_cedula, fecha_solicitud, fecha_inicio, fecha_fin, motivo, aprobado FROM gestion_permisos_asistencia"
	return repository.FetchRows(query)
}

func selectPermisoAsistenciaID(id int) (*sql.Rows, error) {
	query := "SELECT permiso_asistencia_id, obrero_cedula, fecha_solicitud, fecha_inicio, fecha_fin, motivo, aprobado FROM gestion_permisos_asistencia WHERE permiso_asistencia_id = ?"
	return repository.FetchRows(query, id)
}

func insertPermisoAsistencia(data PermisoAsistencia) (sql.Result, error) {
	query := "INSERT INTO gestion_permisos_asistencia (obrero_cedula, fecha_solicitud, fecha_inicio, fecha_fin, motivo, aprobado) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.FechaSolicitud, data.FechaInicio, data.FechaFin, data.Motivo, data.Aprobado)
}

func updatePermisoAsistencia(data PermisoAsistencia) (sql.Result, error) {
	query := "UPDATE gestion_permisos_asistencia SET obrero_cedula = ?, fecha_solicitud = ?, fecha_inicio = ?, fecha_fin = ?, motivo = ?, aprobado = ? WHERE permiso_asistencia_id = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.FechaSolicitud, data.FechaInicio, data.FechaFin, data.Motivo, data.Aprobado, data.PermisoAsistenciaID)
}

func deletePermisoAsistencia(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_permisos_asistencia WHERE permiso_asistencia_id = ?"
	return repository.ExecuteQuery(query, id)
}
