package tipo_eventos

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TipoEventos) error {
    return rows.Scan(
        &dato.IDTipoEventos,
        &dato.NombreTipo,
        &dato.DescripcionTipo,
    )
}

func searchParsedTiposEvento() ([]TipoEventos, error) {
    rows, err := selectTiposEvento()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedTipoEventoID(id int) ([]TipoEventos, error) {
    rows, err := selectTipoEventoID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionTipoEvento(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[TipoEventos](data)
    if err != nil {
        return err
    }

    _, err = insertTipoEvento(dato)
    return err
}

func putTipoEvento(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[TipoEventos](data)
    if err != nil {
        return err
    }

    _, err = updateTipoEvento(dato)
    return err
}

func deletionTipoEvento(id int) error {
    _, err := deleteTipoEvento(id)
    return err
}