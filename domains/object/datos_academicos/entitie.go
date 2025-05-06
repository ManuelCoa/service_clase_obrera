package datos_academicos

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *DatosAcademicos) error {
	return rows.Scan(
		&dato.DatosAcademicosID,
		&dato.ObreroCedula,
		&dato.NivelAcademicoID,
		&dato.InstitucionAcademicaID,
		&dato.TituloAcademicoID,
		&dato.FechaInicio,
		&dato.FechaFinalizacion,
		&dato.CertificacionObtenida,
	)
}

func searchParsedDatosAcademicos() ([]DatosAcademicos, error) {
	rows, err := selectDatosAcademicos()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDatosAcademicosID(id int) ([]DatosAcademicos, error) {
	rows, err := selectDatosAcademicosID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDatosPorObrero(cedula int) ([]DatosAcademicos, error) {
	rows, err := selectDatosPorObrero(cedula)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionDatosAcademicos(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DatosAcademicos](data)
	if err != nil {
		return err
	}

	_, err = insertDatosAcademicos(dato)
	if err != nil {
		return err
	}

	return nil
}

func putDatosAcademicos(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DatosAcademicos](data)
	if err != nil {
		return err
	}

	_, err = updateDatosAcademicos(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionDatosAcademicos(id int) error {
	_, err := deleteDatosAcademicos(id)
	if err != nil {
		return err
	}

	return nil
}
