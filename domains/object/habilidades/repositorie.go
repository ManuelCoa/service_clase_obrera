package habilidades

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectHabilidades() (*sql.Rows, error) {
	query := `
		SELECT habilidad_id, obrero_cedula, tipo_habilidad_id, 
		       nivel_experiencia, certificado, fecha_certificacion,
		       institucion_certificadora
		FROM gestion_habilidades`
	return repository.FetchRows(query)
}

func selectHabilidadID(id int) (*sql.Rows, error) {
	query := `
		SELECT habilidad_id, obrero_cedula, tipo_habilidad_id, 
		       nivel_experiencia, certificado, fecha_certificacion,
		       institucion_certificadora
		FROM gestion_habilidades
		WHERE habilidad_id = ?`
	return repository.FetchRows(query, id)
}

func insertHabilidad(data Habilidades) (sql.Result, error) {
	query := `
		INSERT INTO gestion_habilidades
		(obrero_cedula, tipo_habilidad_id, nivel_experiencia,
		 certificado, fecha_certificacion, institucion_certificadora)
		VALUES (?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.TipoHabilidadID, data.NivelExperiencia,
		data.Certificado, data.FechaCertificacion, data.InstitucionCertificadora)
}

func updateHabilidad(data Habilidades) (sql.Result, error) {
	query := `
		UPDATE gestion_habilidades
		SET obrero_cedula = ?, tipo_habilidad_id = ?, nivel_experiencia = ?,
		    certificado = ?, fecha_certificacion = ?, institucion_certificadora = ?
		WHERE habilidad_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.TipoHabilidadID, data.NivelExperiencia,
		data.Certificado, data.FechaCertificacion, data.InstitucionCertificadora,
		data.HabilidadID)
}

func deleteHabilidad(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_habilidades WHERE habilidad_id = ?"
	return repository.ExecuteQuery(query, id)
}
