package tipo_formacion_extracademica

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposFormacion() (*sql.Rows, error) {
	query := "SELECT id_formacion_extracademica, nombre_tipo, descripcion FROM config_tipo_formacion_extracademica"
	return repository.FetchRows(query)
}

func selectTiposFormacionID(id int) (*sql.Rows, error) {
	query := "SELECT id_formacion_extracademica, nombre_tipo, descripcion FROM config_tipo_formacion_extracademica WHERE id_formacion_extracademica = ?"
	return repository.FetchRows(query, id)
}

func insertTipoFormacion(data TipoFormacionExtracademica) (sql.Result, error) {
	query := "INSERT INTO config_tipo_formacion_extracademica (nombre_tipo, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreTipo, data.Descripcion)
}

func updateTipoFormacion(data TipoFormacionExtracademica) (sql.Result, error) {
	query := "UPDATE config_tipo_formacion_extracademica SET nombre_tipo = ?, descripcion = ? WHERE id_formacion_extracademica = ?"
	return repository.ExecuteQuery(query, data.NombreTipo, data.Descripcion, data.TipoInformacionExtracademica)
}

func deleteTipoFormacion(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_formacion_extracademica WHERE id_formacion_extracademica = ?"
	return repository.ExecuteQuery(query, id)
}
