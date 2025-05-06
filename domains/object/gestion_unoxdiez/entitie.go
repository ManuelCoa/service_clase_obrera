package gestion_unoxdiez

import (
	"claseobrera/domains/entitie"
	"database/sql"
)
func scanData(rows *sql.Rows, dato *GestionUnoxDiez) error {
	return rows.Scan(
		&dato.UnoxDiezID,
		&dato.ObreroCedula,
		&dato.CentroVotacionID,
		&dato.NombresApellidos,
		&dato.Cedula,
		&dato.Telefono,
		&dato.Observacion,
	)
}

func searchParsedGestionUnoxDiez() ([]GestionUnoxDiez, error) {
	rows, err := selectGestionUnoxDiez()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedGestionUnoxDiezID(id int) ([]GestionUnoxDiez, error) {
	rows, err := selectGestionUnoxDiezID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionGestionUnoxDiez(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[GestionUnoxDiez](data)
	if err != nil {
		return err
	}

	_, err = insertGestionUnoxDiez(dato)
	if err != nil {
		return err
	}

	return nil
}

func putGestionUnoxDiez(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[GestionUnoxDiez](data)
	if err != nil {
		return err
	}

	_, err = updateGestionUnoxDiez(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionGestionUnoxDiez(id int) error {
	_, err := deleteGestionUnoxDiez(id)
	if err != nil {
		return err
	}

	return nil
}
