package data_obrero

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectObreros() (*sql.Rows, error) {
	query := `SELECT cedula_obrero, id_institucion, estado_civil, id_piel nombres, apellidos, 
	         fecha_naci, genero, num_telefono, correo_electronico 
	         FROM gestion_datos_obrero`
	return repository.FetchRows(query)//
}
//funcion para consultar por cedula
func selectObreroID(cedula int) (*sql.Rows, error) {
	query := `SELECT cedula_obrero,id_institucion, estado_civil, id_piel, nombres, apellidos, 
	         fecha_naci, genero, num_telefono, correo_electronico 
	         FROM gestion_datos_obrero WHERE cedula_obrero = ?`
	return repository.FetchRows(query, cedula)
}
//funcion para consultar por id de institucion
func selectObreroInstitucion(id_institucion int) (*sql.Rows, error) {
	query := `SELECT 
    d.cedula_obrero,
    d.id_direccion_obrero,
    d.id_estado,
    d.id_municipio,
    d.id_parroquia,
    d.id_ciudad,
    d.id_comuna,
    d.id_consejo,
    d.sector_urbanismo,
    d.direccion,
    d.punto_referencia,
    d.cedula_jefe_calle,
    d.nombre_jefe_calle
FROM 
    gestion_datos_obrero o
JOIN 
    gestion_direccion_obrero d ON o.cedula_obrero = d.cedula_obrero
WHERE 
    o.id_institucion = ?;`
	return repository.FetchRows(query, id_institucion)
}

func insertObrero(data DataObrero) (sql.Result, error) {
	query := `INSERT INTO gestion_datos_obrero 
	         (cedula_obrero, id_institucion, estado_civil, id_piel, nombres, 
	          apellidos, fecha_naci, genero, num_telefono, correo_electronico) 
	         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.CedulaObrero,
		//data.CargoPublicoID,
		//data.CargoOnapreID,
		//data.ResponsabilidadPoliticaID,
		//data.ResponsabilidadComunalID,
		//data.EstructuraPopularID,
		data.InstitucionID,
		//data.ProfesionID,
		data.EstadoCivil,
		data.PielID,
		data.Nombres,
		data.Apellidos,
		data.FechaNacimiento,
		data.Genero,
		//data.TipoTransporte,
		data.NumeroTelefono,
		data.Correro,
	)
}

func updateObrero(data DataObrero) (sql.Result, error) {
	query := `UPDATE gestion_datos_obrero SET id_institucion = ?, estado_civil = ?, id_piel = ?, nombres = ?, apellidos = ?, fecha_naci = ?, genero = ?, 
	         num_telefono = ?, correo_electronico = ? 
	         WHERE cedula_obrero = ?`
	return repository.ExecuteQuery(query,
		//data.CargoPublicoID,
		//data.CargoOnapreID,
		//data.ResponsabilidadPoliticaID,
		//data.ResponsabilidadComunalID,
		//data.EstructuraPopularID,
		data.InstitucionID,
		//data.ProfesionID,
		data.EstadoCivil,
		data.PielID,
		data.Nombres,
		data.Apellidos,
		data.FechaNacimiento,
		data.Genero,
		//data.TipoTransporte,
		data.NumeroTelefono,
		data.Correro,
		data.CedulaObrero,
	)
}

func deleteObrero(cedula int) (sql.Result, error) {
	query := "DELETE FROM gestion_datos_obrero WHERE cedula_obrero = ?"
	return repository.ExecuteQuery(query, cedula)
}