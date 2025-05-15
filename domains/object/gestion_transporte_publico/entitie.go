package transporte_publico

import (
	"claseobrera/domains/entitie"
	"database/sql"
)

func scanData(rows *sql.Rows, dato *TransportePublico) error {
    return rows.Scan(
        &dato.IDTransportePublico,
        &dato.CedulaObrero,
        &dato.RutaTransporte,
        &dato.TiempoEstimado,
    )
}

func searchParsedTransportePublico() ([]TransportePublico, error) {
    rows, err := selectTransportePublico()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func searchParsedTransportePublicoID(id int) ([]TransportePublico, error) {
    rows, err := selectTransportePublicoID(id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    return entitie.ScanRows(rows, scanData)
}

func insertionTransportePublico(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[TransportePublico](data)
    if err != nil {
        return err
    }

    _, err = insertTransportePublico(dato)
    if err != nil {
        return err
    }

    return nil
}

func putTransportePublico(data interface{}) error {
    dato, err := entitie.ParseJSONToStruct[TransportePublico](data)
    if err != nil {
        return err
    }

    _, err = updateTransportePublico(dato)
    if err != nil {
        return err
    }

    return nil
}

func deletionTransportePublico(id int) error {
    _, err := deleteTransportePublico(id)
    if err != nil {
        return err
    }

    return nil
}
