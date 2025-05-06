package dia_feriado

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectDiaFeriado() (*sql.Rows, error) {
	query := "SELECT feriado_id, fecha_dia_feriado, descripcion FROM gestion_dia_feriado"
	return repository.FetchRows(query)
}

func selectDiaFeriadoID(id int) (*sql.Rows, error) {
	query := "SELECT feriado_id, fecha_dia_feriado, descripcion FROM gestion_dia_feriado WHERE feriado_id = ?"
	return repository.FetchRows(query, id)
}

func insertDiaFeriado(data DiaFeriado) (sql.Result, error) {
	query := "INSERT INTO gestion_dia_feriado (fecha_dia_feriado, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.FechaDiaFeriado, data.Descripcion)
}

func updateDiaFeriado(data DiaFeriado) (sql.Result, error) {
	query := "UPDATE gestion_dia_feriado SET fecha_dia_feriado = ?, descripcion = ? WHERE feriado_id = ?"
	return repository.ExecuteQuery(query, data.FechaDiaFeriado, data.Descripcion, data.FeriadoID)
}

func deleteDiaFeriado(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_dia_feriado WHERE feriado_id = ?"
	return repository.ExecuteQuery(query, id)
}
