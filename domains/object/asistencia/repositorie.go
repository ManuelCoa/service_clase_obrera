package asistencia

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectAsistencias() (*sql.Rows, error) {
	query := "SELECT id_asistencia, id_asistencia_biometrica, id_feriado, fecha, hora_entrada, hora_salida FROM gestion_asistencias"
	return repository.FetchRows(query)
}

func selectAsistenciaID(id int) (*sql.Rows, error) {
	query := "SELECT id_asistencia, id_asistencia_biometrica, id_feriado, fecha, hora_entrada, hora_salida FROM gestion_asistencias WHERE id_asistencia = ?"
	return repository.FetchRows(query, id)
}

func insertAsistencia(data Asistencia) (sql.Result, error) {
	query := "INSERT INTO gestion_asistencias (id_asistencia_biometrica, id_feriado, fecha, hora_entrada, hora_salida) VALUES (?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.AsistenciaBiometricaID, data.FeriadoID, data.Fecha, data.HoraEntrada, data.HoraSalida)
}

func updateAsistencia(data Asistencia) (sql.Result, error) {
	query := "UPDATE gestion_asistencias SET id_asistencia_biometrica = ?, id_feriado = ?, fecha = ?, hora_entrada = ?, hora_salida = ? WHERE id_asistencia = ?"
	return repository.ExecuteQuery(query, data.AsistenciaBiometricaID, data.FeriadoID, data.Fecha, data.HoraEntrada, data.HoraSalida, data.AsistenciaID)
}

func deleteAsistencia(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_asistencias WHERE id_asistencia = ?"
	return repository.ExecuteQuery(query, id)
}
