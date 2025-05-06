package organizacion_estructura

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *OrganizacionEstructura) error {
	return rows.Scan(&dato.EstructuraID, &dato.EstructuraNombre, &dato.EstructuraAbrev, &dato.EstructuraDominio)
}

func searchParsedOrganizacionEstructura() ([]OrganizacionEstructura, error) {
	rows, err := selectOrganizacionEstructura()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedOrganizacionEstructuraID(id int) ([]OrganizacionEstructura, error) {
	rows, err := selectOrganizacionEstructuraID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionOrganizacionEstructura(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[OrganizacionEstructura](data)
	if err != nil {
		return err
	}

	_, err = insertOrganizacionEstructura(dato)
	if err != nil {
		return err
	}

	return nil
}

func putOrganizacionEstructura(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[OrganizacionEstructura](data)
	if err != nil {
		return err
	}

	_, err = updateOrganizacionEstructura(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionOrganizacionEstructura(id int) error {
	_, err := deleteOrganizacionEstructura(id)
	if err != nil {
		return err
	}

	return nil
}
