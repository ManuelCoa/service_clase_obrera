package estados

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Estado) error {
	return rows.Scan(&dato.EstadoID, &dato.EstadoNombre)
}

func searchParsedEstado() ([]Estado, error) {
	rows, err := selectEstados()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedEstadoID(id int) ([]Estado, error) {
	rows, err := selectEstadoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionEstado(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Estado](data)
	if err != nil {
		return err
	}

	_, err = insertEstado(dato)
	if err != nil {
		return err
	}

	return nil
}

func putEstado(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Estado](data)
	if err != nil {
		return err
	}

	_, err = updateEstado(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionEstado(id int) error {
	_, err := deleteEstado(id)
	if err != nil {
		return err
	}

	return nil
}
