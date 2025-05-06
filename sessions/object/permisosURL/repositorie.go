package permisosURL

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectPermissionsByRoleID(roleID int) (*sql.Rows, error) {
	query := "SELECT cp.nombre FROM config_user_permisos cp JOIN config_user_perfiles crp ON cp.id_permisos = crp.id_permisos WHERE crp.id_roles = ?"
	return repository.FetchRows(query, roleID)
}
