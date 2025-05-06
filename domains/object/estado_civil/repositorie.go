package estado_civil

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectEstadoCivil() (*sql.Rows, error) {
	query := "SELECT  id_estado_civil, nombre_estado_civil FROM config_estado_civil"
	return repository.FetchRows(query)
}

func selectEstadoCivilID(id int) (*sql.Rows, error) {
	query := "SELECT  id_estado_civil, nombre_estado_civil FROM config_estado_civil WHERE id_estado_civil = ?"
	return repository.FetchRows(query, id)
}

func insertEstadoCivil(data EstadoCivil) (sql.Result, error) {
	query := "INSERT INTO config_estado_civil ( nombre_estado_civil) VALUES (?)"
	return repository.ExecuteQuery(query,data.NombreEdoCivil)
}

func updateEstadoCivil(data EstadoCivil) (sql.Result, error) {
	query := "UPDATE config_estado_civil SET nombre_estado_civil = ? WHERE id_estado_civil = ?"
	return repository.ExecuteQuery(query,data.NombreEdoCivil, data.EstadoCivilID)
}

func deleteEstadoCivil(id int) (sql.Result, error) {
	query := "DELETE FROM config_estado_civil WHERE id_estado_civil = ?"
	return repository.ExecuteQuery(query, id)
}
