package misiones_base

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *MisionesBase) error {
	return rows.Scan(
		&dato.MisionesBaseID,
		&dato.NombreMisiones,
	)
}

func searchParsedMisiones() ([]MisionesBase, error) {
	rows, err := selectMisiones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedMisionesID(id int) ([]MisionesBase, error) {
	rows, err := selectMisionesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionMision(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[MisionesBase](data)
	if err != nil {
		return err
	}

	_, err = insertMision(dato)
	return err
}

func putMision(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[MisionesBase](data)
	if err != nil {
		return err
	}

	_, err = updateMision(dato)
	return err
}

func deletionMision(id int) error {
	_, err := deleteMision(id)
	return err
}