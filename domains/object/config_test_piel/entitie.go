package test_piel

import (
	"database/sql"

	"claseobrera/domains/entitie"
)


func scanData(rows *sql.Rows, dato *TestPiel) error {
	return rows.Scan(&dato.PielID, &dato.NombrePiel)
}

func searchParsedTestPiel() ([]TestPiel, error) {
	rows, err := selectTestPiel()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTestPielID(id int) ([]TestPiel, error) {
	rows, err := selectTestPielID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTestPiel(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TestPiel](data)
	if err != nil {
		return err
	}

	_, err = insertTestPiel(dato)
	if err != nil {
		return err
	}

	return nil
}

func putTestPiel(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TestPiel](data)
	if err != nil {
		return err
	}

	_, err = updateTestPiel(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionTestPiel(id int) error {
	_, err := deleteTestPiel(id)
	if err != nil {
		return err
	}

	return nil
}


