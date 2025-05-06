package consejo_comunal

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *ConsejoComunal) error {
	return rows.Scan(&dato.ConsejoID, &dato.ComunaID, &dato.NombreConsejo, &dato.CodigoConsejo)
}

func searchParsedConsejoComunal() ([]ConsejoComunal, error) {
	rows, err := selectConsejoComunal()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedConsejoComunalID(id int) ([]ConsejoComunal, error) {
	rows, err := selectConsejoComunalID(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionConsejoComunal(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ConsejoComunal](data)
	if err != nil {
		return err
	}

	_, err = insertConsejoComunal(dato)
	if err != nil {
		return err
	}

	return nil
}

func putConsejoComunal(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[ConsejoComunal](data)
	if err != nil {
		return err
	}

	_, err = updateConsejoComunal(dato)
	if err != nil {
		return err
	}

	return nil
}

func deletionConsejoComunal(id int) error {
	_, err := deleteConsejoComunal(id)
	if err != nil {
		return err
	}

	return nil
}
