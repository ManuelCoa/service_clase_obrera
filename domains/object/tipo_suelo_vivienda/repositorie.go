package tipo_suelo_vivienda

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposSuelo() (*sql.Rows, error) {
	query := "SELECT id_tipo_suelo, nombre_tipo_suelo, descripcion_suelo FROM config_tipo_suelo_vivienda"
	return repository.FetchRows(query)
}

func selectTiposSueloID(id int) (*sql.Rows, error) {
	query := "SELECT id_tipo_suelo, nombre_tipo_suelo, descripcion_suelo FROM config_tipo_suelo_vivienda WHERE id_tipo_suelo = ?"
	return repository.FetchRows(query, id)
}

func insertTipoSuelo(data TipoSueloVivienda) (sql.Result, error) {
	query := "INSERT INTO config_tipo_suelo_vivienda (nombre_tipo_suelo, descripcion_suelo) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreTipoSuelo, data.Descripcion)
}

func updateTipoSuelo(data TipoSueloVivienda) (sql.Result, error) {
	query := "UPDATE config_tipo_suelo_vivienda SET nombre_tipo_suelo = ?, descripcion_suelo = ? WHERE id_tipo_suelo = ?"
	return repository.ExecuteQuery(query, data.NombreTipoSuelo, data.Descripcion, data.TipoSueloViviendaID)
}

func deleteTipoSuelo(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_suelo_vivienda WHERE id_tipo_suelo = ?"
	return repository.ExecuteQuery(query, id)
}
