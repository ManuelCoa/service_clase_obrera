package tipo_techo

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoTecho) error {
	return rows.Scan(
		&dato.TipoTechoID,
		&dato.NombreTecho,
		&dato.Descripcion,
	)
}

func searchParsedTiposTecho() ([]TipoTecho, error) {
	rows, err := selectTiposTecho()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposTechoID(id int) ([]TipoTecho, error) {
	rows, err := selectTiposTechoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoTecho(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoTecho](data)
	if err != nil {
		return err
	}

	_, err = insertTipoTecho(dato)
	return err
}

func putTipoTecho(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoTecho](data)
	if err != nil {
		return err
	}

	_, err = updateTipoTecho(dato)
	return err
}

func deletionTipoTecho(id int) error {
	_, err := deleteTipoTecho(id)
	return err
}