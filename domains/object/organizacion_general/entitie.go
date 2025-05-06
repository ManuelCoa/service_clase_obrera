package organizacion_general

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *OrganizacionGeneral) error {
	return rows.Scan(&dato.OrgID, &dato.OrgNombre)
}

func searchParsedOrganizacion() ([]OrganizacionGeneral, error) {
	rows, err := selectOrganizacion()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedOrganizacionID(id int) ([]OrganizacionGeneral, error) {
	rows, err := selectOrganizacionID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionOrganizacion(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[OrganizacionGeneral](data)
	if err != nil {
		return err
	}

	_, err = insertOrganizacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func putOrganizacion(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[OrganizacionGeneral](data)
	if err != nil {
		return err
	}

	_, err = updateOrganizacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionOrganizacion(id int) error {
	_, err := deleteOrganizacion(id)
	if err != nil {
		return err
	}

	return nil
}
