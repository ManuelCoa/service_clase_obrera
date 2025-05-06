package ciudades

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Ciudad) error {
	return rows.Scan(&dato.CiudadID, &dato.EstadoID, &dato.CiudadNombre)
}

func searchParsedCiudad() ([]Ciudad, error) {
	rows, err := selectCiudad()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedCiudadID(id int) ([]Ciudad, error) {
	rows, err := selectCiudadID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionCiudad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Ciudad](data)
	if err != nil {
		return err
	}

	_, err = insertCiudad(dato)
	if err != nil {
		return err
	}

	return nil
}

func putCiudad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Ciudad](data)
	if err != nil {
		return err
	}

	_, err = updateCiudad(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionCiudad(id int) error {
	_, err := deleteCiudad(id)
	if err != nil {
		return err
	}

	return nil
}
