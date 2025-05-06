package justificacion_ausencia

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *JustificacionAusencia) error {
	return rows.Scan(
		&dato.JustificacionID,
		&dato.ObreroCedula,
		&dato.FechaAusencia,
		&dato.Motivo,
		&dato.Justificativo,
	)
}

func searchParsedJustificaciones() ([]JustificacionAusencia, error) {
	rows, err := selectJustificaciones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedJustificacionID(id int) ([]JustificacionAusencia, error) {
	rows, err := selectJustificacionID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionJustificacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[JustificacionAusencia](data)
	if err != nil {
		return err
	}

	_, err = insertJustificacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func putJustificacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[JustificacionAusencia](data)
	if err != nil {
		return err
	}

	_, err = updateJustificacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionJustificacion(id int) error {
	_, err := deleteJustificacion(id)
	if err != nil {
		return err
	}

	return nil
}
