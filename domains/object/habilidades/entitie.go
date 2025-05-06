package habilidades

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Habilidades) error {
	return rows.Scan(
		&dato.HabilidadID,
		&dato.ObreroCedula,
		&dato.TipoHabilidadID,
		&dato.NivelExperiencia,
		&dato.Certificado,
		&dato.FechaCertificacion,
		&dato.InstitucionCertificadora,
	)
}

func searchParsedHabilidades() ([]Habilidades, error) {
	rows, err := selectHabilidades()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedHabilidadID(id int) ([]Habilidades, error) {
	rows, err := selectHabilidadID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionHabilidad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Habilidades](data)
	if err != nil {
		return err
	}

	_, err = insertHabilidad(dato)
	if err != nil {
		return err
	}

	return nil
}

func putHabilidad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Habilidades](data)
	if err != nil {
		return err
	}

	_, err = updateHabilidad(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionHabilidad(id int) error {
	_, err := deleteHabilidad(id)
	if err != nil {
		return err
	}

	return nil
}
