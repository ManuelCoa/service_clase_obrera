package tipo_suelo_vivienda

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoSueloVivienda) error {
	return rows.Scan(
		&dato.TipoSueloViviendaID,
		&dato.NombreTipoSuelo,
		&dato.Descripcion,
	)
}

func searchParsedTiposSuelo() ([]TipoSueloVivienda, error) {
	rows, err := selectTiposSuelo()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposSueloID(id int) ([]TipoSueloVivienda, error) {
	rows, err := selectTiposSueloID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoSuelo(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoSueloVivienda](data)
	if err != nil {
		return err
	}

	_, err = insertTipoSuelo(dato)
	return err
}

func putTipoSuelo(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoSueloVivienda](data)
	if err != nil {
		return err
	}

	_, err = updateTipoSuelo(dato)
	return err
}

func deletionTipoSuelo(id int) error {
	_, err := deleteTipoSuelo(id)
	return err
}