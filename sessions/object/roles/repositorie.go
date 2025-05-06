package roles

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectRoles() (*sql.Rows, error) {
	query := "SELECT id_roles, nombre, descripcion FROM config_user_roles"
	return repository.FetchRows(query)
}

func selectRolesID(id int) (*sql.Rows, error) {
	query := "SELECT id_roles, nombre, descripcion FROM config_user_roles WHERE id_roles = ?"
	return repository.FetchRows(query, id)
}

func insertRoles(data Roles) (sql.Result, error) {
	query := "INSERT INTO config_user_roles (nombre, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.RolesNombre, data.RolesDescripcion)
}

func updateRoles(data Roles) (sql.Result, error) {
	query := "UPDATE config_user_roles SET nombre = ?, descripcion = ? WHERE id_roles = ?"
	return repository.ExecuteQuery(query, data.RolesNombre, data.RolesDescripcion, data.RolesID)
}

func deleteRoles(id int) (sql.Result, error) {
	query := "DELETE FROM config_user_roles WHERE id_roles = ?"
	return repository.ExecuteQuery(query, id)
}
