package antecedentes_medicos

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *AntecedentesMedicos) error {
	return rows.Scan(
		&dato.AntecedentesEnfermedadID,
		&dato.CedulaObrero,
		&dato.TipoEnfermedadID,
		&dato.FechaDiagnostico,
		&dato.DuranteTrabajo,
	)
}

func searchParsedAntecedentesMedicos() ([]AntecedentesMedicos, error) {
	rows, err := selectAntecedentesMedicos()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedAntecedenteMedicoID(id int) ([]AntecedentesMedicos, error) {
	rows, err := selectAntecedenteMedicoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionAntecedenteMedico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[AntecedentesMedicos](data)
	if err != nil {
		return err
	}

	_, err = insertAntecedenteMedico(dato)
	if err != nil {
		return err
	}

	return nil
}

func putAntecedenteMedico(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[AntecedentesMedicos](data)
	if err != nil {
		return err
	}

	_, err = updateAntecedenteMedico(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionAntecedenteMedico(id int) error {
	_, err := deleteAntecedenteMedico(id)
	if err != nil {
		return err
	}

	return nil
}