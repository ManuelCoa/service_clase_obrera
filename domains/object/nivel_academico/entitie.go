package nivel_academico

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *NivelAcademico) error {
	return rows.Scan( &dato.NivelID, &dato.Nombre)
}

func searchParsedNivelAcademico() ([]NivelAcademico, error) {
	rows, err := selectNivelAcademico()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedNivelAcademicoID(id int) ([]NivelAcademico, error) {
	rows, err := selectNivelAcademicoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionNivelAcademico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[NivelAcademico](data)
	if err != nil {
		return err
	}

	_, err = insertNivelAcademico(dato)
	if err != nil {
		return err
	}

	return nil
}

func putNivelAcademico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[NivelAcademico](data)
	if err != nil {
		return err
	}

	_, err = updateNivelAcademico(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionNivelAcademico(id int) error {
	_, err := deleteNivelAcademico(id)
	if err != nil {
		return err
	}

	return nil
}