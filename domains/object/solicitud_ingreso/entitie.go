package solicitud_ingreso

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *SolicitudIngreso) error {
	return rows.Scan(
		&dato.SolicitudIngresoID,
		&dato.InstitucionID,
		&dato.CedulaSolicitante,
		&dato.NombreSolicitante,
		&dato.ApellidoSolicitante,
		&dato.FechaSolicitud,
		&dato.Estatus,
		&dato.Observaciones,
	)
}

func searchParsedSolicitudIngreso() ([]SolicitudIngreso, error) {
	rows, err := selectSolicitudIngreso()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedSolicitudIngresoID(id int) ([]SolicitudIngreso, error) {
	rows, err := selectSolicitudIngresoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionSolicitudIngreso(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[SolicitudIngreso](data)
	if err != nil {
		return err
	}

	_, err = insertSolicitudIngreso(dato)
	if err != nil {
		return err
	}

	return nil
}

func putSolicitudIngreso(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[SolicitudIngreso](data)
	if err != nil {
		return err
	}

	_, err = updateSolicitudIngreso(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionSolicitudIngreso(id int) error {
	_, err := deleteSolicitudIngreso(id)
	if err != nil {
		return err
	}

	return nil
}