package tipo_vivienda

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposVivienda() (*sql.Rows, error) {
	query := "SELECT id_tipo_vivienda, nombre_tipo_vivienda, descripcion_tipo FROM config_tipo_vivienda"
	return repository.FetchRows(query)
}

func selectTiposViviendaID(id int) (*sql.Rows, error) {
	query := "SELECT id_tipo_vivienda, nombre_tipo_vivienda, descripcion_tipo FROM config_tipo_vivienda WHERE id_tipo_vivienda = ?"
	return repository.FetchRows(query, id)
}

func insertTipoVivienda(data TipoVivienda) (sql.Result, error) {
	query := "INSERT INTO config_tipo_vivienda (nombre_tipo_vivienda, descripcion_tipo) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreTipoVivienda, data.Descripcion)
}

func updateTipoVivienda(data TipoVivienda) (sql.Result, error) {
	query := "UPDATE config_tipo_vivienda SET nombre_tipo_vivienda = ?, descripcion_tipo = ? WHERE id_tipo_vivienda = ?"
	return repository.ExecuteQuery(query, data.NombreTipoVivienda, data.Descripcion, data.TipoViviendaID)
}

func deleteTipoVivienda(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_vivienda WHERE id_tipo_vivienda = ?"
	return repository.ExecuteQuery(query, id)
}
