package cargo_publico

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *CargoPublico) error {
	return rows.Scan(&dato.CargoID, &dato.Nombre, &dato.Descripcion)
}

func searchParsedCargo() ([]CargoPublico, error) {
	rows, err := selectCargo()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedCargoID(id int) ([]CargoPublico, error) {
	rows, err := selectCargoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionCargo(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CargoPublico](data)
	if err != nil {
		return err
	}

	_, err = insertCargo(dato)
	if err != nil {
		return err
	}

	return nil
}

func putCargo(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CargoPublico](data)
	if err != nil {
		return err
	}

	_, err = updateCargo(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionCargo(id int) error {
	_, err := deleteCargo(id)
	if err != nil {
		return err
	}

	return nil
}