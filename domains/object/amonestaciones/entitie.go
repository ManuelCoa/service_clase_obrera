package amonestaciones

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Amonestaciones) error {
	return rows.Scan(
		&dato.AmonestacionesID,
		&dato.CedulaObrero,
		&dato.FechaAmonestacion,
		&dato.Motivo,
		&dato.Sancion,
	)
}

func searchParsedAmonestaciones() ([]Amonestaciones, error) {
	rows, err := selectAmonestaciones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedAmonestacionID(id int) ([]Amonestaciones, error) {
	rows, err := selectAmonestacionID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionAmonestacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Amonestaciones](data)
	if err != nil {
		return err
	}

	_, err = insertAmonestacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func putAmonestacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Amonestaciones](data)
	if err != nil {
		return err
	}

	_, err = updateAmonestacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionAmonestacion(id int) error {
	_, err := deleteAmonestacion(id)
	if err != nil {
		return err
	}

	return nil
}