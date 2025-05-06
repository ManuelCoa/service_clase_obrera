package genero

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Genero) error {
	return rows.Scan( &dato.GeneroID, &dato.NombreGenero)
}

func searchParsedGenero() ([]Genero, error) {
	rows, err := selectGenero()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedGeneroID(id int) ([]Genero, error) {
	rows, err := selectGeneroID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionGenero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Genero](data)
	if err != nil {
		return err
	}

	_, err = insertGenero(dato)
	if err != nil {
		return err
	}

	return nil
}

func putGenero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Genero](data)
	if err != nil {
		return err
	}

	_, err = updateGenero(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionGenero(id int) error {
	_, err := deleteGenero(id)
	if err != nil {
		return err
	}

	return nil
}
