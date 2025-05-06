package viviendas

import (
	"claseobrera/domains/entitie"
	"database/sql"
)


func scanData(rows *sql.Rows, dato *Vivienda) error {
	return rows.Scan(
		&dato.ConfigViviendaID,
		&dato.ObreroCedula,
		&dato.TipoViviendaID,
		&dato.TipoSueloID,
		&dato.TipoTechoID,
		&dato.NombreVivienda,
		&dato.NumeroVivienda,
	)
}

func searchParsedVivienda() ([]Vivienda, error) {
	rows, err := selectVivienda()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedViviendaID(id int) ([]Vivienda, error) {
	rows, err := selectViviendaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionVivienda(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Vivienda](data)
	if err != nil {
		return err
	}

	_, err = insertVivienda(dato)
	if err != nil {
		return err
	}

	return nil
}

func putVivienda(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Vivienda](data)
	if err != nil {
		return err
	}

	_, err = updateVivienda(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionVivienda(id int) error {
	_, err := deleteVivienda(id)
	if err != nil {
		return err
	}

	return nil
}
