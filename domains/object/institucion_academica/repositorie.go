package institucion_academica

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectInstitucionesAcademicas() (*sql.Rows, error) {
	query := `
		SELECT institucion_academica_id, estado_id, municipio_id,
		       parroquia_id, ciudad_id, nombre_academia,
		       tipo_institucion, telefono, correo
		FROM gestion_institucion_academica`
	return repository.FetchRows(query)
}

func selectInstitucionAcademicaID(id int) (*sql.Rows, error) {
	query := `
		SELECT institucion_academica_id, estado_id, municipio_id,
		       parroquia_id, ciudad_id, nombre_academia,
		       tipo_institucion, telefono, correo
		FROM gestion_institucion_academica
		WHERE institucion_academica_id = ?`
	return repository.FetchRows(query, id)
}

func insertInstitucionAcademica(data InstitucionAcademica) (sql.Result, error) {
	query := `
		INSERT INTO gestion_institucion_academica
		(estado_id, municipio_id, parroquia_id, ciudad_id,
		 nombre_academia, tipo_institucion, telefono, correo)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.EstadoID, data.MunicipioID, data.ParroquiaID, data.CiudadID,
		data.NombreAcademia, data.TipoInstitucion, data.Telefono, data.Correo)
}

func updateInstitucionAcademica(data InstitucionAcademica) (sql.Result, error) {
	query := `
		UPDATE gestion_institucion_academica
		SET estado_id = ?, municipio_id = ?, parroquia_id = ?, ciudad_id = ?,
		    nombre_academia = ?, tipo_institucion = ?, telefono = ?, correo = ?
		WHERE institucion_academica_id = ?`
	return repository.ExecuteQuery(query,
		data.EstadoID, data.MunicipioID, data.ParroquiaID, data.CiudadID,
		data.NombreAcademia, data.TipoInstitucion, data.Telefono, data.Correo,
		data.InstitucionAcademicaID)
}

func deleteInstitucionAcademica(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_institucion_academica WHERE institucion_academica_id = ?"
	return repository.ExecuteQuery(query, id)
}
