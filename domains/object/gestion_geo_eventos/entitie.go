package geo_eventos

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *GestionGeoEventos) error {
    return rows.Scan(
        &dato.IDGeoEventos,
        &dato.IDEstado,
        &dato.IDMunicipio,
        &dato.IDParroquia,
        &dato.NombreLugar,
    )
}

func searchParsedGeoEventos() ([]GestionGeoEventos, error) {
    rows, err := selectGeoEventos()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedGeoEventoID(id int) ([]GestionGeoEventos, error) {
    rows, err := selectGeoEventoID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionGeoEvento(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[GestionGeoEventos](data)
    if err != nil {
        return err
    }

    _, err = insertGeoEvento(dato)
    return err
}

func putGeoEvento(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[GestionGeoEventos](data)
    if err != nil {
        return err
    }

    _, err = updateGeoEvento(dato)
    return err
}

func deletionGeoEvento(id int) error {
    _, err := deleteGeoEvento(id)
    return err
}
