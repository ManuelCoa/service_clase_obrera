package parroquias

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Parroquia) error {
	return rows.Scan(&dato.ParroquiaID, &dato.MunicipioID, &dato.ParroquiaNombre)
}

func searchParsedParroquia() ([]Parroquia, error) {
	rows, err := selectParroquias()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedParroquiaID(id int) ([]Parroquia, error) {
	rows, err := selectParroquiasID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionParroquia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Parroquia](data)
	if err != nil {
		return err
	}

	_, err = insertParroquia(dato)
	if err != nil {
		return err
	}

	return nil
}

func putParroquia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Parroquia](data)
	if err != nil {
		return err
	}

	_, err = updateParroquia(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionParroquia(id int) error {
	_, err := deleteParroquia(id)
	if err != nil {
		return err
	}

	return nil
}
