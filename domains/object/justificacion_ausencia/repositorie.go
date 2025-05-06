package justificacion_ausencia

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectJustificaciones() (*sql.Rows, error) {
	query := `
		SELECT justificacion_id, obrero_cedula, fecha_ausencia, 
		       motivo, justificativo
		FROM gestion_justificacion_ausencia`
	return repository.FetchRows(query)
}

func selectJustificacionID(id int) (*sql.Rows, error) {
	query := `
		SELECT justificacion_id, obrero_cedula, fecha_ausencia, 
		       motivo, justificativo
		FROM gestion_justificacion_ausencia
		WHERE justificacion_id = ?`
	return repository.FetchRows(query, id)
}

func insertJustificacion(data JustificacionAusencia) (sql.Result, error) {
	query := `
		INSERT INTO gestion_justificacion_ausencia
		(obrero_cedula, fecha_ausencia, motivo, justificativo)
		VALUES (?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.FechaAusencia, data.Motivo, data.Justificativo)
}

func updateJustificacion(data JustificacionAusencia) (sql.Result, error) {
	query := `
		UPDATE gestion_justificacion_ausencia
		SET obrero_cedula = ?, fecha_ausencia = ?, 
		    motivo = ?, justificativo = ?
		WHERE justificacion_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.FechaAusencia, data.Motivo, data.Justificativo,
		data.JustificacionID)
}

func deleteJustificacion(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_justificacion_ausencia WHERE justificacion_id = ?"
	return repository.ExecuteQuery(query, id)
}
