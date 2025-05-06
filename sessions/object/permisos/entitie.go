package permisos

import (
	"database/sql"

	"claseobrera/sessions/entitie"
)

func scanData(rows *sql.Rows, dato *Permisos) error {
	return rows.Scan(&dato.PermisoID, &dato.PermisoNombre, &dato.PermisoDescripcion)
}

func searchParsedPermisos() ([]Permisos, error) {
	rows, err := selectPermisos()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedPermisosID(id int) ([]Permisos, error) {
	rows, err := selectPermisosID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionPermisos(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Permisos](data)
	if err != nil {
		return err
	}

	_, err = insertPermisos(dato)
	if err != nil {
		return err
	}

	return nil
}

func putPermisos(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Permisos](data)
	if err != nil {
		return err
	}

	_, err = updatePermisos(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionPermisos(id int) error {
	_, err := deletePermisos(id)
	if err != nil {
		return err
	}

	return nil
}
