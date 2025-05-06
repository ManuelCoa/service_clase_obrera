package tipo_enfermedad

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposEnfermedad() (*sql.Rows, error) {
	query := "SELECT id_tipo_enfermedad, nombre_enfermedad, descripcion_enfermedad FROM config_tipo_enfermedad"
	return repository.FetchRows(query)
}

func selectTiposEnfermedadID(id int) (*sql.Rows, error) {
	query := "SELECT id_tipo_enfermedad, nombre_enfermedad, descripcion_enfermedad FROM config_tipo_enfermedad WHERE id_tipo_enfermedad = ?"
	return repository.FetchRows(query, id)
}

func insertTipoEnfermedad(data TipoEnfermedad) (sql.Result, error) {
	query := "INSERT INTO config_tipo_enfermedad (nombre_enfermedad, descripcion_enfermedad) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreEnfermedad, data.DescripcionEnfermedad)
}

func updateTipoEnfermedad(data TipoEnfermedad) (sql.Result, error) {
	query := "UPDATE config_tipo_enfermedad SET nombre_enfermedad = ?, descripcion_enfermedad = ? WHERE id_tipo_enfermedad = ?"
	return repository.ExecuteQuery(query, data.NombreEnfermedad, data.DescripcionEnfermedad, data.TipoEnfermedadID)
}

func deleteTipoEnfermedad(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_enfermedad WHERE id_tipo_enfermedad = ?"
	return repository.ExecuteQuery(query, id)
}
