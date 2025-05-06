package gestion_horarios

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *GestionHorario) error {
	return rows.Scan(
		&dato.AsignacionID,
		&dato.ObreroCedula,
		&dato.HorarioID,
		&dato.FechaInicio,
		&dato.FechaFin,
	)
}

func searchParsedGestionHorarios() ([]GestionHorario, error) {
	rows, err := selectGestionHorarios()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedGestionHorarioID(id int) ([]GestionHorario, error) {
	rows, err := selectGestionHorarioID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionGestionHorario(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[GestionHorario](data)
	if err != nil {
		return err
	}

	_, err = insertGestionHorario(dato)
	if err != nil {
		return err
	}

	return nil
}

func putGestionHorario(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[GestionHorario](data)
	if err != nil {
		return err
	}

	_, err = updateGestionHorario(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionGestionHorario(id int) error {
	_, err := deleteGestionHorario(id)
	if err != nil {
		return err
	}

	return nil
}
