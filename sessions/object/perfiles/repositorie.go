package perfiles

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectPerfiles() (*sql.Rows, error) {
	query := "SELECT id_roles, id_permisos FROM config_user_perfiles"
	return repository.FetchRows(query)
}

func insertPerfiles(data Perfiles) (sql.Result, error) {
	query := "INSERT INTO config_user_perfiles (id_roles, id_permisos) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.RolesID, data.PermisosID)
}

func deletePerfiles(rolesID, permisosID int) (sql.Result, error) {
	query := "DELETE FROM config_user_perfiles WHERE id_roles = ? AND id_permisos"
	return repository.ExecuteQuery(query, rolesID, permisosID)
}
