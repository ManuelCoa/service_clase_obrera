package reposo_medico

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *ReposoMedico) error {
	return rows.Scan(
		&dato.ReposoMedicoID,
		&dato.CedulaObrero,
		&dato.FechaInicio,
		&dato.FechaFin,
		&dato.MotivoReposo,
		&dato.MedicoTratante,
	)
}

func searchParsedReposos() ([]ReposoMedico, error) {
	rows, err := selectReposos()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedRepososID(id int) ([]ReposoMedico, error) {
	rows, err := selectRepososID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionReposo(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ReposoMedico](data)
	if err != nil {
		return err
	}

	_, err = insertReposo(dato)
	return err
}

func putReposo(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ReposoMedico](data)
	if err != nil {
		return err
	}

	_, err = updateReposo(dato)
	return err
}

func deletionReposo(id int) error {
	_, err := deleteReposo(id)
	return err
}