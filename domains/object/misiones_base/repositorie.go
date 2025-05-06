package misiones_base

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectMisiones() (*sql.Rows, error) {
	query := "SELECT id_misiones, nombre_misiones FROM config_misiones_base"
	return repository.FetchRows(query)
}

func selectMisionesID(id int) (*sql.Rows, error) {
	query := "SELECT id_misiones, nombre_misiones FROM config_misiones_base WHERE id_misiones = ?"
	return repository.FetchRows(query, id)
}

func insertMision(data MisionesBase) (sql.Result, error) {
	query := "INSERT INTO config_misiones_base (nombre_misiones) VALUES (?)"
	return repository.ExecuteQuery(query, data.NombreMisiones)
}

func updateMision(data MisionesBase) (sql.Result, error) {
	query := "UPDATE config_misiones_base SET nombre_misiones = ? WHERE id_misiones = ?"
	return repository.ExecuteQuery(query, data.NombreMisiones, data.MisionesBaseID)
}

func deleteMision(id int) (sql.Result, error) {
	query := "DELETE FROM config_misiones_base WHERE id_misiones = ?"
	return repository.ExecuteQuery(query, id)
}
