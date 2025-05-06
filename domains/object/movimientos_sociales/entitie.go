package movimientos_sociales

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *MovimientosSociales) error {
	return rows.Scan(&dato.MovimientoID, &dato.NombreMovimiento)
}

func searchParsedMovimientosSociales() ([]MovimientosSociales, error) {
	rows, err := selectMovimientosSociales()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedMovimientosSocialesID(id int) ([]MovimientosSociales, error) {
	rows, err := selectMovimientosSocialesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionMovimientosSociales(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[MovimientosSociales](data)
	if err != nil {
		return err
	}

	_, err = insertMovimientosSociales(dato)
	if err != nil {
		return err
	}

	return nil
}

func putMovimientosSociales(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[MovimientosSociales](data)
	if err != nil {
		return err
	}

	_, err = updateMovimientosSociales(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionMovimientosSociales(id int) error {
	_, err := deleteMovimientosSociales(id)
	if err != nil {
		return err
	}

	return nil
}