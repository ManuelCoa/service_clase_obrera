package tipo_parentesco

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoParentesco) error {
	return rows.Scan(
		&dato.TipoParentescoID,
		&dato.NombreTipoParentesco,
		&dato.Descripcion,
	)
}

func searchParsedTiposParentesco() ([]TipoParentesco, error) {
	rows, err := selectTiposParentesco()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposParentescoID(id int) ([]TipoParentesco, error) {
	rows, err := selectTiposParentescoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoParentesco](data)
	if err != nil {
		return err
	}

	_, err = insertTipoParentesco(dato)
	return err
}

func putTipoParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoParentesco](data)
	if err != nil {
		return err
	}

	_, err = updateTipoParentesco(dato)
	return err
}

func deletionTipoParentesco(id int) error {
	_, err := deleteTipoParentesco(id)
	return err
}