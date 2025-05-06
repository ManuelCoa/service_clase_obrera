package organizacion_instituciones

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Instituciones) error {
	return rows.Scan(&dato.InstitucionID, &dato.OrganizacionID, &dato.JerarquiaID, &dato.EstructuraID, &dato.InstitucionNombre, &dato.InstitucionAbreviatura, &dato.InstitucionDominio)
}

func searchParsedInstituciones() ([]Instituciones, error) {
	rows, err := selectInstituciones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedInstitucionesID(id int) ([]Instituciones, error) {
	rows, err := selectInstitucionesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionInstituciones(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[Instituciones](data)
	if err != nil {
		return err
	}

	_, err = insertInstituciones(dato)
	if err != nil {
		return err
	}

	return nil
}

func putInstituciones(data interface{}) error {

	dato, err := entitie.ParseJSONToStruct[Instituciones](data)
	if err != nil {
		return err
	}

	_, err = updateInstituciones(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionInstituciones(id int) error {
	_, err := deleteInstituciones(id)
	if err != nil {
		return err
	}

	return nil
}
