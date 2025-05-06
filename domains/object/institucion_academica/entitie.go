package institucion_academica

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *InstitucionAcademica) error {
	return rows.Scan(
		&dato.InstitucionAcademicaID,
		&dato.EstadoID,
		&dato.MunicipioID,
		&dato.ParroquiaID,
		&dato.CiudadID,
		&dato.NombreAcademia,
		&dato.TipoInstitucion,
		&dato.Telefono,
		&dato.Correo,
	)
}

func searchParsedInstitucionesAcademicas() ([]InstitucionAcademica, error) {
	rows, err := selectInstitucionesAcademicas()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedInstitucionAcademicaID(id int) ([]InstitucionAcademica, error) {
	rows, err := selectInstitucionAcademicaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionInstitucionAcademica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[InstitucionAcademica](data)
	if err != nil {
		return err
	}

	_, err = insertInstitucionAcademica(dato)
	if err != nil {
		return err
	}

	return nil
}

func putInstitucionAcademica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[InstitucionAcademica](data)
	if err != nil {
		return err
	}

	_, err = updateInstitucionAcademica(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionInstitucionAcademica(id int) error {
	_, err := deleteInstitucionAcademica(id)
	if err != nil {
		return err
	}

	return nil
}
