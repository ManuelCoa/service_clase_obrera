package registro_talento_humano

import (
	"database/sql"

	"claseobrera/domains/entitie"
)


func scanData(rows *sql.Rows, dato *RegistroTalentoHumano) error {
	return rows.Scan(
		&dato.RegistroID,
		&dato.ObreroCedula,
		&dato.CargoID,
		&dato.FechaIngreso,
	)
}

func searchParsedRegistroTalentoHumano() ([]RegistroTalentoHumano, error) {
	rows, err := selectRegistroTalentoHumano()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedRegistroTalentoHumanoID(id int) ([]RegistroTalentoHumano, error) {
	rows, err := selectRegistroTalentoHumanoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionRegistroTalentoHumano(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[RegistroTalentoHumano](data)
	if err != nil {
		return err
	}

	_, err = insertRegistroTalentoHumano(dato)
	if err != nil {
		return err
	}

	return nil
}

func putRegistroTalentoHumano(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[RegistroTalentoHumano](data)
	if err != nil {
		return err
	}

	_, err = updateRegistroTalentoHumano(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionRegistroTalentoHumano(id int) error {
	_, err := deleteRegistroTalentoHumano(id)
	if err != nil {
		return err
	}

	return nil
}
