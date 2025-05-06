package asistencia_biometrica

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *AsistenciaBiometrica) error {
	return rows.Scan(
		&dato.AsistenciaBiometricaID,
		&dato.CedulaObrero,
		&dato.IdentificacionBiometrica,
		&dato.FechaRegistro,
		&dato.Activa,
	)
}

func searchParsedAsistenciasBiometricas() ([]AsistenciaBiometrica, error) {
	rows, err := selectAsistenciasBiometricas()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedAsistenciaBiometricaID(id int) ([]AsistenciaBiometrica, error) {
	rows, err := selectAsistenciaBiometricaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionAsistenciaBiometrica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[AsistenciaBiometrica](data)
	if err != nil {
		return err
	}

	_, err = insertAsistenciaBiometrica(dato)
	if err != nil {
		return err
	}

	return nil
}

func putAsistenciaBiometrica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[AsistenciaBiometrica](data)
	if err != nil {
		return err
	}

	_, err = updateAsistenciaBiometrica(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionAsistenciaBiometrica(id int) error {
	_, err := deleteAsistenciaBiometrica(id)
	if err != nil {
		return err
	}

	return nil
}
