package data_obrero

import (
	"claseobrera/domains/repository"
	"database/sql"
	"fmt"
)

func selectObreros() (*sql.Rows, error) {
	query := `SELECT cedula_obrero, id_institucion, id_estado_civil, id_piel, nombres, apellidos, 
	         fecha_naci, id_genero, num_telefono, correo_electronico
	         FROM gestion_datos_obrero`
	return repository.FetchRows(query)
}
//funcion para consultar por cedula
func selectObreroID(cedula int) (*sql.Rows, error) {
	query := `SELECT cedula_obrero,id_institucion, id_estado_civil, id_piel, nombres, apellidos, 
	         fecha_naci, id_genero, num_telefono, correo_electronico 
	         FROM gestion_datos_obrero WHERE cedula_obrero = ?`
	return repository.FetchRows(query, cedula)
}
//funcion para consultar por id de institucion
func selectObreroInstitucion(id_institucion int) (*sql.Rows, error) {
	query := `SELECT cedula_obrero,  id_institucion, id_estado_civil, id_piel, nombres, apellidos, 
	         fecha_naci, id_genero, num_telefono, correo_electronico 
	         FROM gestion_datos_obrero WHERE id_institucion = ?`
	return repository.FetchRows(query, id_institucion)
}

func insertObrero(data DataObrero) (sql.Result, error) {
    // Imprimir el valor de data
    fmt.Printf("Datos a insertar: %+v\n", data)

    query := `INSERT INTO gestion_datos_obrero 
             (cedula_obrero, id_institucion, id_estado_civil, id_piel, nombres, 
              apellidos, fecha_naci, id_genero, num_telefono, correo_electronico) 
             VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

    return repository.ExecuteQuery(query,
        data.CedulaObrero,
        data.InstitucionID,
        data.EstadoCivil,
        data.PielID,
        data.Nombres,
        data.Apellidos,
        data.FechaNacimiento,
        data.Genero,
        data.NumeroTelefono,
        data.Correo,
    )
}

func updateObrero(data DataObrero) (sql.Result, error) {
	query := `UPDATE gestion_datos_obrero SET id_institucion = ?, id_estado_civil = ?, id_piel = ?, nombres = ?, apellidos = ?, fecha_naci = ?, id_genero = ?, 
	         num_telefono = ?, correo_electronico = ? 
	         WHERE cedula_obrero = ?`
	return repository.ExecuteQuery(query,
		data.InstitucionID,
		data.EstadoCivil,
		data.PielID,
		data.Nombres,
		data.Apellidos,
		data.FechaNacimiento,
		data.Genero,
		data.NumeroTelefono,
		data.Correo,
		data.CedulaObrero,
	)
}

func deleteObrero(cedula int) (sql.Result, error) {
	query := "DELETE FROM gestion_datos_obrero WHERE cedula_obrero = ?"
	return repository.ExecuteQuery(query, cedula)
}