package marca_vehiculo

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *MarcaVehiculo) error {
	return rows.Scan(
		&dato.MarcaVehiculoID,
		&dato.NombreMarca,
	)
}

func searchParsedMarcas() ([]MarcaVehiculo, error) {
	rows, err := selectMarcas()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedMarcasID(id int) ([]MarcaVehiculo, error) {
	rows, err := selectMarcasID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionMarca(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[MarcaVehiculo](data)
	if err != nil {
		return err
	}

	_, err = insertMarca(dato)
	return err
}

func putMarca(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[MarcaVehiculo](data)
	if err != nil {
		return err
	}

	_, err = updateMarca(dato)
	return err
}

func deletionMarca(id int) error {
	_, err := deleteMarca(id)
	return err
}