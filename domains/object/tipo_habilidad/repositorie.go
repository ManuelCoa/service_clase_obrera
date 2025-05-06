package tipo_habilidad

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposHabilidad() (*sql.Rows, error) {
	query := "SELECT id_tipo_habilidad, nombre_habilidad, descripcion_habilidad FROM config_tipo_habilidad"
	return repository.FetchRows(query)
}

func selectTiposHabilidadID(id int) (*sql.Rows, error) {
	query := "SELECT id_tipo_habilidad, nombre_habilidad, descripcion_habilidad FROM config_tipo_habilidad WHERE id_tipo_habilidad = ?"
	return repository.FetchRows(query, id)
}

func insertTipoHabilidad(data TipoHabilidad) (sql.Result, error) {
	query := "INSERT INTO config_tipo_habilidad (nombre_habilidad, descripcion_habilidad) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreHabilidad, data.DescripcionHabilidad)
}

func updateTipoHabilidad(data TipoHabilidad) (sql.Result, error) {
	query := "UPDATE config_tipo_habilidad SET nombre_habilidad = ?, descripcion_habilidad = ? WHERE id_tipo_habilidad = ?"
	return repository.ExecuteQuery(query, data.NombreHabilidad, data.DescripcionHabilidad, data.TipoHabilidadID)
}

func deleteTipoHabilidad(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_habilidad WHERE id_tipo_habilidad = ?"
	return repository.ExecuteQuery(query, id)
}
