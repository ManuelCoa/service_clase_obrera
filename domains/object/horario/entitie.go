package horario

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *Horario) error {
	return rows.Scan(
		&dato.HorariioID,
		&dato.ConfigHorario,
		&dato.HoraEntrada,
		&dato.HoraSalida,
		&dato.Descripcion,
	)
}

func searchParsedHorarios() ([]Horario, error) {
	rows, err := selectHorarios()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedHorariosID(id int) ([]Horario, error) {
	rows, err := selectHorariosID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionHorario(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Horario](data)
	if err != nil {
		return err
	}

	_, err = insertHorario(dato)
	return err
}

func putHorario(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Horario](data)
	if err != nil {
		return err
	}

	_, err = updateHorario(dato)
	return err
}

func deletionHorario(id int) error {
	_, err := deleteHorario(id)
	return err
}