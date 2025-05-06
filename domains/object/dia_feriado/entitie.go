package dia_feriado

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *DiaFeriado) error {
	return rows.Scan(&dato.FeriadoID, &dato.FechaDiaFeriado, &dato.Descripcion)
}

func searchParsedDiaFeriado() ([]DiaFeriado, error) {
	rows, err := selectDiaFeriado()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedDiaFeriadoID(id int) ([]DiaFeriado, error) {
	rows, err := selectDiaFeriadoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionDiaFeriado(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DiaFeriado](data)
	if err != nil {
		return err
	}

	_, err = insertDiaFeriado(dato)
	if err != nil {
		return err
	}

	return nil
}

func putDiaFeriado(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DiaFeriado](data)
	if err != nil {
		return err
	}

	_, err = updateDiaFeriado(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionDiaFeriado(id int) error {
	_, err := deleteDiaFeriado(id)
	if err != nil {
		return err
	}

	return nil
}