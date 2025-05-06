package responsabilidades_comunales

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *ResponsabilidadComunal) error {
	return rows.Scan(
		&dato.ResponsabilidadComunalID,
		&dato.NombreResponsabilidad,
		&dato.DescripcionResponsabilidad,
	)
}

func searchParsedResponsabilidades() ([]ResponsabilidadComunal, error) {
	rows, err := selectResponsabilidades()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedResponsabilidadesID(id int) ([]ResponsabilidadComunal, error) {
	rows, err := selectResponsabilidadesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionResponsabilidad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ResponsabilidadComunal](data)
	if err != nil {
		return err
	}

	_, err = insertResponsabilidad(dato)
	return err
}

func putResponsabilidad(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ResponsabilidadComunal](data)
	if err != nil {
		return err
	}

	_, err = updateResponsabilidad(dato)
	return err
}

func deletionResponsabilidad(id int) error {
	_, err := deleteResponsabilidad(id)
	return err
}