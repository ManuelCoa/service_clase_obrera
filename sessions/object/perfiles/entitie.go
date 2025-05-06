package perfiles

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanData(rows *sql.Rows, dato *Perfiles) error {
	return rows.Scan(&dato.RolesID, &dato.PermisosID)
}

func searchParsedPerfiles() ([]Perfiles, error) {
	rows, err := selectPerfiles()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionPerfiles(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Perfiles](data)
	if err != nil {
		return err
	}

	_, err = insertPerfiles(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionPerfiles(rolesID, permisosID int) error {
	_, err := deletePerfiles(rolesID, permisosID)
	if err != nil {
		return err
	}

	return nil
}
