package gestionextracademica

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectGestionExtracademica() (*sql.Rows, error) {
	query := `
		SELECT info_extracademica_id, obrero_cedula, tipo_formacion_extracademica_id,
		       certificacion_extracademica_id, institucion_extracademica_id,
		       fecha_inicio, fecha_finalizacion
		FROM gestion_extracademica`
	return repository.FetchRows(query)
}

func selectGestionExtracademicaID(id int) (*sql.Rows, error) {
	query := `
		SELECT info_extracademica_id, obrero_cedula, tipo_formacion_extracademica_id,
		       certificacion_extracademica_id, institucion_extracademica_id,
		       fecha_inicio, fecha_finalizacion
		FROM gestion_extracademica
		WHERE info_extracademica_id = ?`
	return repository.FetchRows(query, id)
}

func insertGestionExtracademica(data GestionExtracademica) (sql.Result, error) {
	query := `
		INSERT INTO gestion_extracademica
		(obrero_cedula, tipo_formacion_extracademica_id, certificacion_extracademica_id,
		 institucion_extracademica_id, fecha_inicio, fecha_finalizacion)
		VALUES (?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.TipoFormacionExtracademicaID, data.CertificacionExtracademicaID,
		data.InstitucionExtracademicaID, data.FechaInicio, data.FechaFinalizacion)
}

func updateGestionExtracademica(data GestionExtracademica) (sql.Result, error) {
	query := `
		UPDATE gestion_extracademica
		SET obrero_cedula = ?, tipo_formacion_extracademica_id = ?,
		    certificacion_extracademica_id = ?, institucion_extracademica_id = ?,
		    fecha_inicio = ?, fecha_finalizacion = ?
		WHERE info_extracademica_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.TipoFormacionExtracademicaID, data.CertificacionExtracademicaID,
		data.InstitucionExtracademicaID, data.FechaInicio, data.FechaFinalizacion,
		data.InfoExtracademicaID)
}

func deleteGestionExtracademica(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_extracademica WHERE info_extracademica_id = ?"
	return repository.ExecuteQuery(query, id)
}