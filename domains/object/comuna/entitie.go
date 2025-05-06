package comuna

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *Comuna) error {
	return rows.Scan(&dato.ComunaID, &dato.CentroVotacionID, &dato.EstadoID, &dato.MunicipioID, &dato.ParroquiaID, &dato.NombreComuna, &dato.CodigoCircuitoComuna)
}

func searchParsedComuna() ([]Comuna, error) {
	rows, err := selectComuna()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedComunaID(id int) ([]Comuna, error) {
	rows, err := selectComunaID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionComuna(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Comuna](data)
	if err != nil {
		return err
	}

	_, err = insertComuna(dato)
	if err != nil {
		return err
	}

	return nil
}

func putComuna(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[Comuna](data)
	if err != nil {
		return err
	}

	_, err = updateComuna(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionComuna(id int) error {
	_, err := deleteComuna(id)
	if err != nil {
		return err
	}

	return nil
}