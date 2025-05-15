package direccion_parentesco

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *DireccionParentesco) error {
    return rows.Scan(
        &dato.IDDireccionParentesco,
        &dato.CedulaParentesco,
        &dato.EstadoID,
        &dato.MunicipioID,
        &dato.ParroquiaID,
        &dato.CiudadID,
        &dato.SectorUrbanismo,
        &dato.Direccion,
        &dato.PuntoReferencia,
    )
}

func searchParsedDireccionParentesco() ([]DireccionParentesco, error) {
    rows, err := selectDireccionParentesco()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedDireccionParentescoID(id int) ([]DireccionParentesco, error) {
    rows, err := selectDireccionParentescoID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionDireccionParentesco(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[DireccionParentesco](data)
    if err != nil {
        return err
    }

    _, err = insertDireccionParentesco(dato)
    if err != nil {
        return err
    }

    return nil
}

func putDireccionParentesco(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[DireccionParentesco](data)
    if err != nil {
        return err
    }

    _, err = updateDireccionParentesco(dato)
    if err != nil {
        return err
    }

    return nil
}

func deletionDireccionParentesco(id int) error {
    _, err := deleteDireccionParentesco(id)
    if err != nil {
        return err
    }

    return nil
}