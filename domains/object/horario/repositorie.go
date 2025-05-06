package horario

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectHorarios() (*sql.Rows, error) {
	query := "SELECT id_horario, config_horario, hora_entrada, hora_salida, descripcion FROM horario"
	return repository.FetchRows(query)
}

func selectHorariosID(id int) (*sql.Rows, error) {
	query := "SELECT id_horario, config_horario, hora_entrada, hora_salida, descripcion FROM horario WHERE id_horario = ?"
	return repository.FetchRows(query, id)
}

func insertHorario(data Horario) (sql.Result, error) {
	query := "INSERT INTO horario (config_horario, hora_entrada, hora_salida, descripcion) VALUES (?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ConfigHorario, data.HoraEntrada, data.HoraSalida, data.Descripcion)
}

func updateHorario(data Horario) (sql.Result, error) {
	query := "UPDATE horario SET config_horario = ?, hora_entrada = ?, hora_salida = ?, descripcion = ? WHERE id_horario = ?"
	return repository.ExecuteQuery(query, data.ConfigHorario, data.HoraEntrada, data.HoraSalida, data.Descripcion, data.HorariioID)
}

func deleteHorario(id int) (sql.Result, error) {
	query := "DELETE FROM horario WHERE id_horario = ?"
	return repository.ExecuteQuery(query, id)
}
