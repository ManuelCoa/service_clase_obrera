package gestionextracademica

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *GestionExtracademica) error {
	return rows.Scan(
		&dato.InfoExtracademicaID,
		&dato.ObreroCedula,
		&dato.TipoFormacionExtracademicaID,
		&dato.CertificacionExtracademicaID,
		&dato.InstitucionExtracademicaID,
		&dato.FechaInicio,
		&dato.FechaFinalizacion,
	)
}

func searchParsedGestionExtracademica() ([]GestionExtracademica, error) {
	rows, err := selectGestionExtracademica()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedGestionExtracademicaID(id int) ([]GestionExtracademica, error) {
	rows, err := selectGestionExtracademicaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionGestionExtracademica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[GestionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = insertGestionExtracademica(dato)
	if err != nil {
		return err
	}

	return nil
}

func putGestionExtracademica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[GestionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = updateGestionExtracademica(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionGestionExtracademica(id int) error {
	_, err := deleteGestionExtracademica(id)
	if err != nil {
		return err
	}

	return nil
}