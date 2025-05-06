package profesiones

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Profesiones) error {
	return rows.Scan(&dato.ProfesionID, &dato.NombreProfesion)
}

func searchParsedProfesiones() ([]Profesiones, error) {
	rows, err := selectProfesiones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedProfesionesID(id int) ([]Profesiones, error) {
	rows, err := selectProfesionesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionProfesiones(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Profesiones](data)
	if err != nil {
		return err
	}

	_, err = insertProfesiones(dato)
	if err != nil {
		return err
	}

	return nil
}

func putProfesiones(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Profesiones](data)
	if err != nil {
		return err
	}

	_, err = updateProfesiones(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionProfesiones(id int) error {
	_, err := deleteProfesiones(id)
	if err != nil {
		return err
	}

	return nil
}