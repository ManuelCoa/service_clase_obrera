package eventos_populares

import (
	"database/sql"

	"claseobrera/domains/entitie"
)

func scanData(rows *sql.Rows, dato *GestionEventosPopulares) error {
    return rows.Scan(
        &dato.IDEvento,
        &dato.IDGeoEventos,
        &dato.IDTipoEventos,
        &dato.IDInstitucion,
        &dato.NombreEvento,
        &dato.CuotaPersonas,
        &dato.FechaEvento,
        &dato.HoraEvento,
        &dato.DescripcionEvento,
        &dato.EventoActivo,
    )
}

func searchParsedEventos() ([]GestionEventosPopulares, error) {
    rows, err := selectEventos()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedEventoID(id int) ([]GestionEventosPopulares, error) {
    rows, err := selectEventoID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionEvento(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[GestionEventosPopulares](data)
    if err != nil {
        return err
    }

    _, err = insertEvento(dato)
    return err
}

func putEvento(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[GestionEventosPopulares](data)
    if err != nil {
        return err
    }

    _, err = updateEvento(dato)
    return err
}

func deletionEvento(id int) error {
    _, err := deleteEvento(id)
    return err
}
