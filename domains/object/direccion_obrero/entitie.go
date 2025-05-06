package direccion_obrero

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *DireccionObrero) error {
	return rows.Scan(
		&dato.DireccionObreroID,
		&dato.ObreroCedula,
		&dato.EstadoID,
		&dato.MunicipioID,
		&dato.ParroquiaID,
		&dato.CiudadID,
		&dato.SectorUrbanismo,
		&dato.Direccion,
		&dato.PuntoReferencia,
	)
}

func searchParsedDireccionObrero() ([]DireccionObrero, error) {
	rows, err := selectDireccionObrero()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDireccionObreroID(id int) ([]DireccionObrero, error) {
	rows, err := selectDireccionObreroID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionDireccionObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DireccionObrero](data)
	if err != nil {
		return err
	}

	_, err = insertDireccionObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func putDireccionObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DireccionObrero](data)
	if err != nil {
		return err
	}

	_, err = updateDireccionObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionDireccionObrero(id int) error {
	_, err := deleteDireccionObrero(id)
	if err != nil {
		return err
	}

	return nil
}