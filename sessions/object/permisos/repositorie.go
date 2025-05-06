package permisos

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectPermisos() (*sql.Rows, error) {
	query := "SELECT id_permisos, nombre, descripcion FROM config_user_permisos"
	return repository.FetchRows(query)
}

func selectPermisosID(id int) (*sql.Rows, error) {
	query := "SELECT id_permisos, nombre, descripcion FROM config_user_permisos WHERE id_permisos = ?"
	return repository.FetchRows(query, id)
}

func insertPermisos(data Permisos) (sql.Result, error) {
	query := "INSERT INTO config_user_permisos (nombre, descripcion) VALUES (?)"
	return repository.ExecuteQuery(query, data.PermisoNombre, data.PermisoDescripcion)
}

func updatePermisos(data Permisos) (sql.Result, error) {
	query := "UPDATE config_user_permisos SET nombre = ?, descripcion = ? WHERE id_permisos = ?"
	return repository.ExecuteQuery(query, data.PermisoNombre, data.PermisoDescripcion, data.PermisoID)
}

func deletePermisos(id int) (sql.Result, error) {
	query := "DELETE FROM config_user_permisos WHERE id_permisos = ?"
	return repository.ExecuteQuery(query, id)
}
