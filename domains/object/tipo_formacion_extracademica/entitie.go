package tipo_formacion_extracademica

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoFormacionExtracademica) error {
	return rows.Scan(
		&dato.TipoInformacionExtracademica,
		&dato.NombreTipo,
		&dato.Descripcion,
	)
}

func searchParsedTiposFormacion() ([]TipoFormacionExtracademica, error) {
	rows, err := selectTiposFormacion()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedTiposFormacionID(id int) ([]TipoFormacionExtracademica, error) {
	rows, err := selectTiposFormacionID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionTipoFormacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoFormacionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = insertTipoFormacion(dato)
	return err
}

func putTipoFormacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[TipoFormacionExtracademica](data)
	if err != nil {
		return err
	}

	_, err = updateTipoFormacion(dato)
	return err
}

func deletionTipoFormacion(id int) error {
	_, err := deleteTipoFormacion(id)
	return err
}