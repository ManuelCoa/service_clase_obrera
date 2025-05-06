package asistencia

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Asistencia) error {
	return rows.Scan(
		&dato.AsistenciaID,
		&dato.AsistenciaBiometricaID,
		&dato.FeriadoID,
		&dato.Fecha,
		&dato.HoraEntrada,
		&dato.HoraSalida,
	)
}

func searchParsedAsistencias() ([]Asistencia, error) {
	rows, err := selectAsistencias()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedAsistenciaID(id int) ([]Asistencia, error) {
	rows, err := selectAsistenciaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionAsistencia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Asistencia](data)
	if err != nil {
		return err
	}

	_, err = insertAsistencia(dato)
	if err != nil {
		return err
	}

	return nil
}

func putAsistencia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Asistencia](data)
	if err != nil {
		return err
	}

	_, err = updateAsistencia(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionAsistencia(id int) error {
	_, err := deleteAsistencia(id)
	if err != nil {
		return err
	}

	return nil
}
