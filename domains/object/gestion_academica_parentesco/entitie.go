package academica_parentesco

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *AcademicaParentesco) error {
	return rows.Scan(
		&dato.DatosAcademicosID,
		&dato.CedulaParentesco,
		&dato.NivelAcademicoID,
		&dato.TituloAcademicoID,
		&dato.InstitucionAcademicaID,
		&dato.FechaInicio,
		&dato.FechaFin,
		&dato.Certificado,
	)
}

func searchParsedAcademicasParentesco() ([]AcademicaParentesco, error) {
	rows, err := selectAcademicasParentesco()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedAcademicaParentescoID(id int) ([]AcademicaParentesco, error) {
	rows, err := selectAcademicaParentescoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionAcademicaParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[AcademicaParentesco](data)
	if err != nil {
		return err
	}

	_, err = insertAcademicaParentesco(dato)
	if err != nil {
		return err
	}

	return nil
}

func putAcademicaParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[AcademicaParentesco](data)
	if err != nil {
		return err
	}

	_, err = updateAcademicaParentesco(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionAcademicaParentesco(id int) error {
	_, err := deleteAcademicaParentesco(id)
	if err != nil {
		return err
	}

	return nil
}