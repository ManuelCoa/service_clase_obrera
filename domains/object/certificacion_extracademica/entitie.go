package certificacion_extracademica

import (
	"database/sql"

	"claseobrera/domains/entitie"
)



func scanData(rows *sql.Rows, dato *CertificacionExtracademica) error {
	return rows.Scan(&dato.CertificacionID, &dato.NumeroCertificacion, &dato.FechaEgreso)
}

func searchParsedCertificaciones() ([]CertificacionExtracademica, error) {
	rows, err := selectCertificaciones()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedCertificacionID(id int) ([]CertificacionExtracademica, error) {
	rows, err := selectCertificacionID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionCertificacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CertificacionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = insertCertificacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func putCertificacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CertificacionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = updateCertificacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionCertificacion(id int) error {
	_, err := deleteCertificacion(id)
	if err != nil {
		return err
	}

	return nil
}
