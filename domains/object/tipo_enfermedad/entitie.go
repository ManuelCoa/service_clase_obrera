package tipo_enfermedad

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoEnfermedad) error {
	return rows.Scan(
		&dato.TipoEnfermedadID,
		&dato.NombreEnfermedad,
		&dato.DescripcionEnfermedad,
	)
}

func searchParsedTiposEnfermedad() ([]TipoEnfermedad, error) {
	rows, err := selectTiposEnfermedad()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposEnfermedadID(id int) ([]TipoEnfermedad, error) {
	rows, err := selectTiposEnfermedadID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoEnfermedad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoEnfermedad](data)
	if err != nil {
		return err
	}

	_, err = insertTipoEnfermedad(dato)
	return err
}

func putTipoEnfermedad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoEnfermedad](data)
	if err != nil {
		return err
	}

	_, err = updateTipoEnfermedad(dato)
	return err
}

func deletionTipoEnfermedad(id int) error {
	_, err := deleteTipoEnfermedad(id)
	return err
}