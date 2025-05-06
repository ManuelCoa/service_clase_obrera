package bioquimicos_obrero

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *BioquimicosObrero) error {
	return rows.Scan(
		&dato.BioquimicosObreroID,
		&dato.CedulaObrero,
		&dato.FechaAnalisis,
		&dato.NivelGlucosa,
		&dato.NivelColesterol,
		&dato.NivelTrigliceridos,
		&dato.PresionArterial,
	)
}

func searchParsedBioquimicosObreros() ([]BioquimicosObrero, error) {
	rows, err := selectBioquimicosObreros()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedBioquimicosObreroID(id int) ([]BioquimicosObrero, error) {
	rows, err := selectBioquimicosObreroID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionBioquimicosObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[BioquimicosObrero](data)
	if err != nil {
		return err
	}

	_, err = insertBioquimicosObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func putBioquimicosObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[BioquimicosObrero](data)
	if err != nil {
		return err
	}

	_, err = updateBioquimicosObrero(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionBioquimicosObrero(id int) error {
	_, err := deleteBioquimicosObrero(id)
	if err != nil {
		return err
	}

	return nil
}
