package parentesco

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Parentesco) error {
	return rows.Scan(&dato.ParentescoID, &dato.NombreParentesco)
}

func searchParsedParentesco() ([]Parentesco, error) {
	rows, err := selectParentesco()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedParentescoID(id int) ([]Parentesco, error) {
	rows, err := selectParentescoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Parentesco](data)
	if err != nil {
		return err
	}

	_, err = insertParentesco(dato)
	if err != nil {
		return err
	}

	return nil
}

func putParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Parentesco](data)
	if err != nil {
		return err
	}

	_, err = updateParentesco(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionParentesco(id int) error {
	_, err := deleteParentesco(id)
	if err != nil {
		return err
	}

	return nil
}