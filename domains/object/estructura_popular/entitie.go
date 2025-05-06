package estructura_popular

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *EstructuraPopular) error {
	return rows.Scan(&dato.EstructuraID, &dato.NombreEstructura, &dato.Descripcion)
}

func searchParsedEstructuraPopular() ([]EstructuraPopular, error) {
	rows, err := selectEstructuraPopular()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedEstructuraPopularID(id int) ([]EstructuraPopular, error) {
	rows, err := selectEstructuraPopularID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionEstructuraPopular(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[EstructuraPopular](data)
	if err != nil {
		return err
	}

	_, err = insertEstructuraPopular(dato)
	if err != nil {
		return err
	}

	return nil
}

func putEstructuraPopular(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[EstructuraPopular](data)
	if err != nil {
		return err
	}

	_, err = updateEstructuraPopular(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionEstructuraPopular(id int) error {
	_, err := deleteEstructuraPopular(id)
	if err != nil {
		return err
	}

	return nil
}