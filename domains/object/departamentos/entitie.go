package departamentos

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *Departamentos) error {
	return rows.Scan(
		&dato.DepartamentoID,
		&dato.InstitucionID,
		&dato.NombreDepartamento,
		&dato.DescripcionDepartamento,
	)
}

func searchParsedDepartamentos() ([]Departamentos, error) {
	rows, err := selectDepartamentos()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDepartamentosID(id int) ([]Departamentos, error) {
	rows, err := selectDepartamentosID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionDepartamentos(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Departamentos](data)
	if err != nil {
		return err
	}

	_, err = insertDepartamentos(dato)
	return err
}

func putDepartamentos(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Departamentos](data)
	if err != nil {
		return err
	}

	_, err = updateDepartamentos(dato)
	return err
}

func deletionDepartamentos(id int) error {
	_, err := deleteDepartamentos(id)
	return err
}