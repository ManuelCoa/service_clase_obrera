package tipo_habilidad

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoHabilidad) error {
	return rows.Scan(
		&dato.TipoHabilidadID,
		&dato.NombreHabilidad,
		&dato.DescripcionHabilidad,
	)
}

func searchParsedTiposHabilidad() ([]TipoHabilidad, error) {
	rows, err := selectTiposHabilidad()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposHabilidadID(id int) ([]TipoHabilidad, error) {
	rows, err := selectTiposHabilidadID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoHabilidad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoHabilidad](data)
	if err != nil {
		return err
	}

	_, err = insertTipoHabilidad(dato)
	return err
}

func putTipoHabilidad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoHabilidad](data)
	if err != nil {
		return err
	}

	_, err = updateTipoHabilidad(dato)
	return err
}

func deletionTipoHabilidad(id int) error {
	_, err := deleteTipoHabilidad(id)
	return err
}