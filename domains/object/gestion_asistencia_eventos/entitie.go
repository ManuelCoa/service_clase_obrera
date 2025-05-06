package asistencia_eventos

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *GestionAsistenciaEventos) error {
    return rows.Scan(
        &dato.IDAsistenciaEventos,
        &dato.IDEvento,
        &dato.CedulaObrero,
        &dato.IDInstitucion,
        &dato.Participacion,
        &dato.Observacion,
    )
}

func searchParsedAsistencias() ([]GestionAsistenciaEventos, error) {
    rows, err := selectAsistencias()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedAsistenciaID(id int) ([]GestionAsistenciaEventos, error) {
    rows, err := selectAsistenciaID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionAsistencia(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[GestionAsistenciaEventos](data)
    if err != nil {
        return err
    }

    _, err = insertAsistencia(dato)
    return err
}

func putAsistencia(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[GestionAsistenciaEventos](data)
    if err != nil {
        return err
    }

    _, err = updateAsistencia(dato)
    return err
}

func deletionAsistencia(id int) error {
    _, err := deleteAsistencia(id)
    return err
}
