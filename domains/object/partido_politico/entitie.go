package partido_politico

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *PartidoPolitico) error {
	return rows.Scan(&dato.PartidoID, &dato.NombrePartido)
}

func searchParsedPartidoPolitico() ([]PartidoPolitico, error) {
	rows, err := selectPartidoPolitico()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedPartidoPoliticoID(id int) ([]PartidoPolitico, error) {
	rows, err := selectPartidoPoliticoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionPartidoPolitico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[PartidoPolitico](data)
	if err != nil {
		return err
	}

	_, err = insertPartidoPolitico(dato)
	if err != nil {
		return err
	}

	return nil
}

func putPartidoPolitico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[PartidoPolitico](data)
	if err != nil {
		return err
	}

	_, err = updatePartidoPolitico(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionPartidoPolitico(id int) error {
	_, err := deletePartidoPolitico(id)
	if err != nil {
		return err
	}

	return nil
}