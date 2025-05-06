package constancia_trabajo

import (
	"claseobrera/domains/entitie"
	"database/sql"
)


func scanData(rows *sql.Rows, dato *ConstanciaTrabajo) error {
	return rows.Scan(&dato.ConstanciaID, &dato.ObreroCedula, &dato.InstitucionID, &dato.FechaEmision, &dato.Desempeño, &dato.Observaciones)
}

func searchParsedConstancias() ([]ConstanciaTrabajo, error) {
	rows, err := selectConstancias()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedConstanciaID(id int) ([]ConstanciaTrabajo, error) {
	rows, err := selectConstanciaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionConstancia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ConstanciaTrabajo](data)
	if err != nil {
		return err
	}

	_, err = insertConstancia(dato)
	if err != nil {
		return err
	}

	return nil
}

func putConstancia(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ConstanciaTrabajo](data)
	if err != nil {
		return err
	}

	_, err = updateConstancia(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionConstancia(id int) error {
	_, err := deleteConstancia(id)
	if err != nil {
		return err
	}

	return nil
}