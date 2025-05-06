package cargos_onapre

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *CargosOnapre) error {
	return rows.Scan(&dato.CargoOnapreID, &dato.Nombre, &dato.Descripcion)
}

func searchParsedCargoOnapre() ([]CargosOnapre, error) {
	rows, err := selectCargoOnapre()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedCargoOnapreID(id int) ([]CargosOnapre, error) {
	rows, err := selectCargoOnapreID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionCargoOnapre(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CargosOnapre](data)
	if err != nil {
		return err
	}

	_, err = insertCargoOnapre(dato)
	if err != nil {
		return err
	}

	return nil
}

func putCargoOnapre(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CargosOnapre](data)
	if err != nil {
		return err
	}

	_, err = updateCargoOnapre(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionCargoOnapre(id int) error {
	_, err := deleteCargoOnapre(id)
	if err != nil {
		return err
	}

	return nil
}
