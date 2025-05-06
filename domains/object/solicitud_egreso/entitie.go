package solicitud_egreso

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *SolicitudEgreso) error {
	return rows.Scan(
		&dato.SolicitudEgresoID,
		&dato.ObreroCedula,
		&dato.FechaSolicitud,
		&dato.Motivo,
		&dato.Estatus,
		&dato.Observaciones,
	)
}

func searchParsedSolicitudEgreso() ([]SolicitudEgreso, error) {
	rows, err := selectSolicitudEgreso()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedSolicitudEgresoID(id int) ([]SolicitudEgreso, error) {
	rows, err := selectSolicitudEgresoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionSolicitudEgreso(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[SolicitudEgreso](data)
	if err != nil {
		return err
	}

	_, err = insertSolicitudEgreso(dato)
	if err != nil {
		return err
	}

	return nil
}

func putSolicitudEgreso(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[SolicitudEgreso](data)
	if err != nil {
		return err
	}

	_, err = updateSolicitudEgreso(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionSolicitudEgreso(id int) error {
	_, err := deleteSolicitudEgreso(id)
	if err != nil {
		return err
	}

	return nil
}
