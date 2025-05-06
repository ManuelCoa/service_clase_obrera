package centro_votacion

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *CentroVotacion) error {
	return rows.Scan(&dato.CentroVotacionID, &dato.EstadoID, &dato.MunicipioID, &dato.ParroquiaID, &dato.CodigoCentro, &dato.NombreCentro, &dato.DireccionCentro, &dato.CodigoNuevo)
}

func searchParsedCentroVotacion() ([]CentroVotacion, error) {
	rows, err := selectCentroVotacion()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedCentroVotacionID(id int) ([]CentroVotacion, error) {
	rows, err := selectCentroVotacionID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionCentroVotacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CentroVotacion](data)
	if err != nil {
		return err
	}

	_, err = insertCentroVotacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func putCentroVotacion(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[CentroVotacion](data)
	if err != nil {
		return err
	}

	_, err = updateCentroVotacion(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionCentroVotacion(id int) error {
	_, err := deleteCentroVotacion(id)
	if err != nil {
		return err
	}

	return nil
}