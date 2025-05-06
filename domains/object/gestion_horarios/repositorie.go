package gestion_horarios

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectGestionHorarios() (*sql.Rows, error) {
	query := `
		SELECT asignacion_id, obrero_cedula, horario_id, 
		       fecha_inicio, fecha_fin
		FROM gestion_horario`
	return repository.FetchRows(query)
}

func selectGestionHorarioID(id int) (*sql.Rows, error) {
	query := `
		SELECT asignacion_id, obrero_cedula, horario_id, 
		       fecha_inicio, fecha_fin
		FROM gestion_horario
		WHERE asignacion_id = ?`
	return repository.FetchRows(query, id)
}

func insertGestionHorario(data GestionHorario) (sql.Result, error) {
	query := `
		INSERT INTO gestion_horario
		(obrero_cedula, horario_id, fecha_inicio, fecha_fin)
		VALUES (?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.HorarioID, data.FechaInicio, data.FechaFin)
}

func updateGestionHorario(data GestionHorario) (sql.Result, error) {
	query := `
		UPDATE gestion_horario
		SET obrero_cedula = ?, horario_id = ?, 
		    fecha_inicio = ?, fecha_fin = ?
		WHERE asignacion_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.HorarioID, data.FechaInicio, data.FechaFin,
		data.AsignacionID)
}

func deleteGestionHorario(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_horario WHERE asignacion_id = ?"
	return repository.ExecuteQuery(query, id)
}
