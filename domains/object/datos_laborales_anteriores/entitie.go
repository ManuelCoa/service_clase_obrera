package datos_laborales

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *DatosLaboralesAnteriores) error {
	return rows.Scan(
		&dato.LaboralesDatosID,
		&dato.ObreroCedula,
		&dato.InstitucionID,
		&dato.CargoID,
		&dato.DepartamentoID,
		&dato.FechaInicio,
		&dato.FechaFin,
		&dato.MotivoEgreso,
	)
}

func searchParsedDatosLaborales() ([]DatosLaboralesAnteriores, error) {
	rows, err := selectDatosLaborales()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDatosLaboralesID(id int) ([]DatosLaboralesAnteriores, error) {
	rows, err := selectDatosLaboralesID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDatosPorObrero(cedula int) ([]DatosLaboralesAnteriores, error) {
	rows, err := selectDatosPorObrero(cedula)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionDatosLaborales(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DatosLaboralesAnteriores](data)
	if err != nil {
		return err
	}

	_, err = insertDatosLaborales(dato)
	if err != nil {
		return err
	}

	return nil
}

func putDatosLaborales(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DatosLaboralesAnteriores](data)
	if err != nil {
		return err
	}

	_, err = updateDatosLaborales(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionDatosLaborales(id int) error {
	_, err := deleteDatosLaborales(id)
	if err != nil {
		return err
	}

	return nil
}
