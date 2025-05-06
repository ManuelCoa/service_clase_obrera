package roles

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanData(rows *sql.Rows, dato *Roles) error {
	return rows.Scan(&dato.RolesID, &dato.RolesNombre, &dato.RolesDescripcion)
}

func searchParsedRoles() ([]Roles, error) {
	rows, err := selectRoles()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedRolesID(id int) ([]Roles, error) {
	rows, err := selectRolesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionRoles(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Roles](data)
	if err != nil {
		return err
	}

	_, err = insertRoles(dato)
	if err != nil {
		return err
	}

	return nil
}

func putRoles(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Roles](data)
	if err != nil {
		return err
	}

	_, err = updateRoles(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionRoles(id int) error {
	_, err := deleteRoles(id)
	if err != nil {
		return err
	}

	return nil
}
