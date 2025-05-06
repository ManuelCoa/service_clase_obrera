package institucion_extracademica

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *InstitucionExtracademica) error {
	return rows.Scan(
		&dato.InstitucionID,
		&dato.NombreInstitucion,
		&dato.Telefono,
		&dato.Correo,
	)
}

func searchParsedInstituciones() ([]InstitucionExtracademica, error) {
	rows, err := selectInstituciones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedInstitucionesID(id int) ([]InstitucionExtracademica, error) {
	rows, err := selectInstitucionesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionInstitucion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[InstitucionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = insertInstitucion(dato)
	return err
}

func putInstitucion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[InstitucionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = updateInstitucion(dato)
	return err
}

func deletionInstitucion(id int) error {
	_, err := deleteInstitucion(id)
	return err
}