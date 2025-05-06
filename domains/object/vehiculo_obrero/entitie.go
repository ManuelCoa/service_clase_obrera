package vehiculo_obrero

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *VehiculoObrero) error {
	return rows.Scan(
		&dato.VehiculoID,
		&dato.MarcaVehiculoID,
		&dato.ObreroCedula,
		&dato.Placa,
		&dato.Modelo,
		&dato.Color,
		&dato.NivelGasolina,
	)
}

func searchParsedVehiculoObrero() ([]VehiculoObrero, error) {
	rows, err := selectVehiculoObrero()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedVehiculoObreroID(id int) ([]VehiculoObrero, error) {
	rows, err := selectVehiculoObreroID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionVehiculoObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[VehiculoObrero](data)
	if err != nil {
		return err
	}

	_, err = insertVehiculoObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func putVehiculoObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[VehiculoObrero](data)
	if err != nil {
		return err
	}

	_, err = updateVehiculoObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionVehiculoObrero(id int) error {
	_, err := deleteVehiculoObrero(id)
	if err != nil {
		return err
	}

	return nil
}
