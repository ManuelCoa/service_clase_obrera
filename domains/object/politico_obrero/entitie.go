package politico_obrero

import (
	"claseobrera/domains/entitie"
	"database/sql"
)
func scanData(rows *sql.Rows, dato *PoliticoObrero) error {
	return rows.Scan(
		&dato.PoliticoObreroID,
		&dato.ObreroCedula,
		&dato.FuerzaPoliticaID,
		&dato.PPID,
		&dato.MSID,
		&dato.MisionesID,
		&dato.RegistroUnoxDiez,
		&dato.Observacion,
	)
}

func searchParsedPoliticoObrero() ([]PoliticoObrero, error) {
	rows, err := selectPoliticoObrero()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedPoliticoObreroID(id int) ([]PoliticoObrero, error) {
	rows, err := selectPoliticoObreroID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionPoliticoObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[PoliticoObrero](data)
	if err != nil {
		return err
	}

	_, err = insertPoliticoObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func putPoliticoObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[PoliticoObrero](data)
	if err != nil {
		return err
	}

	_, err = updatePoliticoObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionPoliticoObrero(id int) error {
	_, err := deletePoliticoObrero(id)
	if err != nil {
		return err
	}

	return nil
}
