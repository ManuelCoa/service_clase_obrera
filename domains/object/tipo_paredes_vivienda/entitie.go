package paredes_vivienda

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoParedesVivienda) error {
    return rows.Scan(
        &dato.IDTipoPared,
        &dato.NombrePared,
        &dato.DescripcionPared,
    )
}

func searchParsedTiposPared() ([]TipoParedesVivienda, error) {
    rows, err := selectTiposPared()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedTipoParedID(id int) ([]TipoParedesVivienda, error) {
    rows, err := selectTipoParedID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionTipoPared(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[TipoParedesVivienda](data)
    if err != nil {
        return err
    }

    _, err = insertTipoPared(dato)
    return err
}

func putTipoPared(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[TipoParedesVivienda](data)
    if err != nil {
        return err
    }

    _, err = updateTipoPared(dato)
    return err
}

func deletionTipoPared(id int) error {
    _, err := deleteTipoPared(id)
    return err
}