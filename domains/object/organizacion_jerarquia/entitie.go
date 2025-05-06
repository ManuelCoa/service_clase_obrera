package organizacion_jerarquia

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *OrganizacionJerarquia) error {
	return rows.Scan(&dato.JerarquiaID, &dato.OrganizacionID, &dato.JerarquiaNombre)
}

func searchParsedOrganizacionJerarquia() ([]OrganizacionJerarquia, error) {
	rows, err := selectOrganizacionJerarquia()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedOrganizacionJerarquiaID(id int) ([]OrganizacionJerarquia, error) {
	rows, err := selectOrganizacionJerarquiaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionOrganizacionJerarquia(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[OrganizacionJerarquia](data)
	if err != nil {
		return err
	}

	_, err = insertOrganizacionJerarquia(dato)
	if err != nil {
		return err
	}

	return nil
}

func putOrganizacionJerarquia(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[OrganizacionJerarquia](data)
	if err != nil {
		return err
	}

	_, err = updateOrganizacionJerarquia(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionOrganizacionJerarquia(id int) error {
	_, err := deleteOrganizacionJerarquia(id)
	if err != nil {
		return err
	}

	return nil
}
