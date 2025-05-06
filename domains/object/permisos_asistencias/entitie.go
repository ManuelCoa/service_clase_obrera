package permisos_asistencias

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *PermisoAsistencia) error {
	return rows.Scan(
		&dato.PermisoAsistenciaID,
		&dato.ObreroCedula,
		&dato.FechaSolicitud,
		&dato.FechaInicio,
		&dato.FechaFin,
		&dato.Motivo,
		&dato.Aprobado,
	)
}

func searchParsedPermisoAsistencia() ([]PermisoAsistencia, error) {
	rows, err := selectPermisoAsistencia()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedPermisoAsistenciaID(id int) ([]PermisoAsistencia, error) {
	rows, err := selectPermisoAsistenciaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionPermisoAsistencia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[PermisoAsistencia](data)
	if err != nil {
		return err
	}

	_, err = insertPermisoAsistencia(dato)
	if err != nil {
		return err
	}

	return nil
}

func putPermisoAsistencia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[PermisoAsistencia](data)
	if err != nil {
		return err
	}

	_, err = updatePermisoAsistencia(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionPermisoAsistencia(id int) error {
	_, err := deletePermisoAsistencia(id)
	if err != nil {
		return err
	}

	return nil
}

