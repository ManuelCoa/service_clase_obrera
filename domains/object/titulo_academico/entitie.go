package titulo_academico

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *TituloAcademico) error {
	return rows.Scan(
		&dato.TituloAcademicoID,
		&dato.CodigoTitulo,
		&dato.FechaEgreso,
	)
}

func searchParsedTituloAcademico() ([]TituloAcademico, error) {
	rows, err := selectTituloAcademico()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTituloAcademicoID(id int) ([]TituloAcademico, error) {
	rows, err := selectTituloAcademicoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTituloAcademico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TituloAcademico](data)
	if err != nil {
		return err
	}

	_, err = insertTituloAcademico(dato)
	if err != nil {
		return err
	}

	return nil
}

func putTituloAcademico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TituloAcademico](data)
	if err != nil {
		return err
	}

	_, err = updateTituloAcademico(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionTituloAcademico(id int) error {
	_, err := deleteTituloAcademico(id)
	if err != nil {
		return err
	}

	return nil
}
