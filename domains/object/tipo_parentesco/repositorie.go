package tipo_parentesco

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposParentesco() (*sql.Rows, error) {
	query := "SELECT id_tipo_parentesco, nombre_tipo_parentesco, descripcion FROM config_tipo_parentesco"
	return repository.FetchRows(query)
}

func selectTiposParentescoID(id int) (*sql.Rows, error) {
	query := "SELECT id_tipo_parentesco, nombre_tipo_parentesco, descripcion FROM config_tipo_parentesco WHERE id_tipo_parentesco = ?"
	return repository.FetchRows(query, id)
}

func insertTipoParentesco(data TipoParentesco) (sql.Result, error) {
	query := "INSERT INTO config_tipo_parentesco (nombre_tipo_parentesco, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreTipoParentesco, data.Descripcion)
}

func updateTipoParentesco(data TipoParentesco) (sql.Result, error) {
	query := "UPDATE config_tipo_parentesco SET nombre_tipo_parentesco = ?, descripcion = ? WHERE id_tipo_parentesco = ?"
	return repository.ExecuteQuery(query, data.NombreTipoParentesco, data.Descripcion, data.TipoParentescoID)
}

func deleteTipoParentesco(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_parentesco WHERE id_tipo_parentesco = ?"
	return repository.ExecuteQuery(query, id)
}
