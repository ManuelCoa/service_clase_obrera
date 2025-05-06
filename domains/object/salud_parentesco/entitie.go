package salud_parentesco

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *SaludParentesco) error {
	return rows.Scan(
		&dato.SaludParentescoID,
		&dato.ParentescoCedula,
		&dato.TipoEnfermedadID,
		&dato.Tratamiento,
	)
}

func searchParsedSaludParentesco() ([]SaludParentesco, error) {
	rows, err := selectSaludParentesco()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedSaludParentescoID(id int) ([]SaludParentesco, error) {
	rows, err := selectSaludParentescoID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionSaludParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[SaludParentesco](data)
	if err != nil {
		return err
	}

	_, err = insertSaludParentesco(dato)
	if err != nil {
		return err
	}

	return nil
}

func putSaludParentesco(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[SaludParentesco](data)
	if err != nil {
		return err
	}

	_, err = updateSaludParentesco(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionSaludParentesco(id int) error {
	_, err := deleteSaludParentesco(id)
	if err != nil {
		return err
	}

	return nil
}