package municipios

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanDato(rows *sql.Rows, dato *Municipio) error {
	return rows.Scan(&dato.MunicipioID, &dato.EstadoID, &dato.MunicipioNombre)
}

func searchParsedMunicipios() ([]Municipio, error) {
	rows, err := selectMunicipios()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func searchParsedMunicipioID(id int) ([]Municipio, error) {
	rows, err := selectMunicipioID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanDato)
}

func insertionMunicipio(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Municipio](data)
	if err != nil {
		return err
	}

	_, err = insertMunicipio(dato)
	if err != nil {
		return err
	}

	return nil
}

func putMunicipio(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Municipio](data)
	if err != nil {
		return err
	}

	_, err = updateMunicipio(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionMunicipio(id int) error {
	_, err := deleteMunicipio(id)
	if err != nil {
		return err
	}

	return nil
}
