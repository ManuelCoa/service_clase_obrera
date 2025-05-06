package estado_civil

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *EstadoCivil) error {
	return rows.Scan( &dato.EstadoCivilID, &dato.NombreEdoCivil)
}

func searchParsedEstadoCivil() ([]EstadoCivil, error) {
	rows, err := selectEstadoCivil()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedEstadoCivilID(id int) ([]EstadoCivil, error) {
	rows, err := selectEstadoCivilID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionEstadoCivil(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[EstadoCivil](data)
	if err != nil {
		return err
	}

	_, err = insertEstadoCivil(dato)
	if err != nil {
		return err
	}

	return nil
}

func putEstadoCivil(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[EstadoCivil](data)
	if err != nil {
		return err
	}

	_, err = updateEstadoCivil(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionEstadoCivil(id int) error {
	_, err := deleteEstadoCivil(id)
	if err != nil {
		return err
	}

	return nil
}
