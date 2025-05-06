package responsabilidad_politica

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *ResponsabilidadPolitica) error {
	return rows.Scan(&dato.ResponsabilidadID, &dato.Nombre, &dato.Descripcion)
}

func searchParsedResponsabilidadPolitica() ([]ResponsabilidadPolitica, error) {
	rows, err := selectResponsabilidadPolitica()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedResponsabilidadPoliticaID(id int) ([]ResponsabilidadPolitica, error) {
	rows, err := selectResponsabilidadPoliticaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionResponsabilidadPolitica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ResponsabilidadPolitica](data)
	if err != nil {
		return err
	}

	_, err = insertResponsabilidadPolitica(dato)
	if err != nil {
		return err
	}

	return nil
}

func putResponsabilidadPolitica(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ResponsabilidadPolitica](data)
	if err != nil {
		return err
	}

	_, err = updateResponsabilidadPolitica(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionResponsabilidadPolitica(id int) error {
	_, err := deleteResponsabilidadPolitica(id)
	if err != nil {
		return err
	}

	return nil
}