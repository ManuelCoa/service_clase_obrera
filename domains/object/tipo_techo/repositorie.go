package tipo_techo

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposTecho() (*sql.Rows, error) {
	query := "SELECT id_tipo_techo, nombre_techo, descripcion_techo FROM config_tipo_techo"
	return repository.FetchRows(query)
}

func selectTiposTechoID(id int) (*sql.Rows, error) {
	query := "SELECT id_tipo_techo, nombre_techo, descripcion_techo FROM config_tipo_techo WHERE id_tipo_techo = ?"
	return repository.FetchRows(query, id)
}

func insertTipoTecho(data TipoTecho) (sql.Result, error) {
	query := "INSERT INTO config_tipo_techo (nombre_techo, descripcion_techo) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreTecho, data.Descripcion)
}

func updateTipoTecho(data TipoTecho) (sql.Result, error) {
	query := "UPDATE config_tipo_techo SET nombre_techo = ?, descripcion_techo = ? WHERE id_tipo_techo = ?"
	return repository.ExecuteQuery(query, data.NombreTecho, data.Descripcion, data.TipoTechoID)
}

func deleteTipoTecho(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_techo WHERE id_tipo_techo = ?"
	return repository.ExecuteQuery(query, id)
}
