package fuerza_politica

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *FuerzaPolitica) error {
	return rows.Scan(
		&dato.FuerzaPoliticaID,
		&dato.NombreFuerza,
	)
}

func searchParsedFuerzas() ([]FuerzaPolitica, error) {
	rows, err := selectFuerzas()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedFuerzasID(id int) ([]FuerzaPolitica, error) {
	rows, err := selectFuerzasID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionFuerza(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[FuerzaPolitica](data)
	if err != nil {
		return err
	}

	_, err = insertFuerza(dato)
	return err
}

func putFuerza(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[FuerzaPolitica](data)
	if err != nil {
		return err
	}

	_, err = updateFuerza(dato)
	return err
}

func deletionFuerza(id int) error {
	_, err := deleteFuerza(id)
	return err
}