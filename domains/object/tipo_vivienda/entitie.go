package tipo_vivienda

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoVivienda) error {
	return rows.Scan(
		&dato.TipoViviendaID,
		&dato.NombreTipoVivienda,
		&dato.Descripcion,
	)
}

func searchParsedTiposVivienda() ([]TipoVivienda, error) {
	rows, err := selectTiposVivienda()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposViviendaID(id int) ([]TipoVivienda, error) {
	rows, err := selectTiposViviendaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoVivienda(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoVivienda](data)
	if err != nil {
		return err
	}

	_, err = insertTipoVivienda(dato)
	return err
}

func putTipoVivienda(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoVivienda](data)
	if err != nil {
		return err
	}

	_, err = updateTipoVivienda(dato)
	return err
}

func deletionTipoVivienda(id int) error {
	_, err := deleteTipoVivienda(id)
	return err
}