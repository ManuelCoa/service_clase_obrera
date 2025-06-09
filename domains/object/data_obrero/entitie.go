package data_obrero

import (
	"claseobrera/domains/entitie"
	"database/sql"
	//"claseobrera/domains/object/direccion_obrero"
)

func scanData(rows *sql.Rows, dato *DataObrero) error {
	return rows.Scan(
		&dato.CedulaObrero,
		&dato.InstitucionID,
		&dato.EstadoCivil,
		&dato.PielID,
		&dato.Nombres,
		&dato.Apellidos,
		&dato.FechaNacimiento,
		&dato.Genero,
		&dato.NumeroTelefono,
		&dato.Correo,
	)
}

func searchParsedObreros() ([]DataObrero, error) {
	rows, err := selectObreros()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func searchParsedObreroID(cedula int) ([]DataObrero, error) {
	rows, err := selectObreroID(cedula)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}
//nueva funcion para filtarar por id 
func searchParsedObreroIDInstitucion(id_institucion int) ([]DataObrero, error) {
	rows, err := selectObreroInstitucion(id_institucion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return entitie.ScanRows(rows, scanData)
}

func insertionObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DataObrero](data)
	if err != nil {
		return err
	}

	_, err = insertObrero(dato)
	return err
}

func putObrero(data interface{}) error {
	dato, err := entitie.ParseJSONToStruct[DataObrero](data)
	if err != nil {
		return err
	}

	_, err = updateObrero(dato)
	return err
}

func deletionObrero(cedula int) error {
	_, err := deleteObrero(cedula)
	return err
}