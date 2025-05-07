package data_obrero

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectObreros() (*sql.Rows, error) {
	query := `SELECT cedula_obrero, id_cargo_publico, id_cargo_onapre, responsabilidad_politica, 
	         responsabilidad_comunal, id_estructura_popular, id_institucion, id_profesion, estado_civil, nombres, apellidos, 
	         fecha_naci, genero, tipo_transporte, num_telefono, correo_electronico 
	         FROM data_obrero`
	return repository.FetchRows(query)
}

func selectObreroID(cedula int) (*sql.Rows, error) {
	query := `SELECT cedula_obrero, id_cargo_publico, id_cargo_onapre, responsabilidad_politica, 
	         responsabilidad_comunal, id_estructura_popular, id_institucion, id_profesion, estado_civil, nombres, apellidos, 
	         fecha_naci, genero, tipo_transporte, num_telefono, correo_electronico 
	         FROM data_obrero WHERE cedula_obrero = ?`
	return repository.FetchRows(query, cedula)
}

func insertObrero(data DataObrero) (sql.Result, error) {
	query := `INSERT INTO data_obrero 
	         (cedula_obrero, id_cargo_publico, id_cargo_onapre, responsabilidad_politica, 
	          responsabilidad_comunal, id_estructura_popular, id_institucion, id_profesion, estado_civil, nombres, 
	          apellidos, fecha_naci, genero, tipo_transporte, num_telefono, correo_electronico) 
	         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.CedulaObrero,
		data.CargoPublicoID,
		data.CargoOnapreID,
		data.ResponsabilidadPoliticaID,
		data.ResponsabilidadComunalID,
		data.EstructuraPopularID,
		data.InstitucionID,
		data.ProfesionID,
		data.EstadoCivil,
		data.Nombres,
		data.Apellidos,
		data.FechaNacimiento,
		data.Genero,
		data.TipoTransporte,
		data.NumeroTelefono,
		data.Correro,
	)
}

func updateObrero(data DataObrero) (sql.Result, error) {
	query := `UPDATE data_obrero SET 
	         id_cargo_publico = ?, id_cargo_onapre = ?, responsabilidad_politica = ?, 
	         responsabilidad_comunal = ?, id_estructura_popular = ?,  id_institucion = ?, id_profesion = ?, 
	         estado_civil = ?, nombres = ?, apellidos = ?, fecha_naci = ?, genero = ?, 
	         tipo_transporte = ?, num_telefono = ?, correo_electronico = ? 
	         WHERE cedula_obrero = ?`
	return repository.ExecuteQuery(query,
		data.CargoPublicoID,
		data.CargoOnapreID,
		data.ResponsabilidadPoliticaID,
		data.ResponsabilidadComunalID,
		data.EstructuraPopularID,
		data.InstitucionID,
		data.ProfesionID,
		data.EstadoCivil,
		data.Nombres,
		data.Apellidos,
		data.FechaNacimiento,
		data.Genero,
		data.TipoTransporte,
		data.NumeroTelefono,
		data.Correro,
		data.CedulaObrero,
	)
}

func deleteObrero(cedula int) (sql.Result, error) {
	query := "DELETE FROM data_obrero WHERE cedula_obrero = ?"
	return repository.ExecuteQuery(query, cedula)
}