package permisosURL

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanData(rows *sql.Rows, dato *PermisosURL) error {
	return rows.Scan(&dato.URL)
}

func searchPermissionsByRoleID(id int) ([]PermisosURL, error) {
	rows, err := selectPermissionsByRoleID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}
