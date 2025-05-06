package datos_academicos

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectDatosAcademicos() (*sql.Rows, error) {
	query := `
		SELECT dato_academico_id, obrero_cedula, nivel_academico_id, 
		       institucion_academica_id, titulo_academico_id, fecha_inicio, 
		       fecha_finalizacion, certificacion_obtenida 
		FROM gestion_dato_academico`
	return repository.FetchRows(query)
}

func selectDatosAcademicosID(id int) (*sql.Rows, error) {
	query := `
		SELECT dato_academico_id, obrero_cedula, nivel_academico_id, 
		       institucion_academica_id, titulo_academico_id, fecha_inicio, 
		       fecha_finalizacion, certificacion_obtenida 
		FROM gestion_dato_academico 
		WHERE dato_academico_id = ?`
	return repository.FetchRows(query, id)
}

func selectDatosPorObrero(cedula int) (*sql.Rows, error) {
	query := `
		SELECT dato_academico_id, obrero_cedula, nivel_academico_id, 
		       institucion_academica_id, titulo_academico_id, fecha_inicio, 
		       fecha_finalizacion, certificacion_obtenida 
		FROM gestion_dato_academico 
		WHERE obrero_cedula = ?`
	return repository.FetchRows(query, cedula)
}

func insertDatosAcademicos(data DatosAcademicos) (sql.Result, error) {
	query := `
		INSERT INTO gestion_dato_academico 
		(obrero_cedula, nivel_academico_id, institucion_academica_id, 
		 titulo_academico_id, fecha_inicio, fecha_finalizacion, certificacion_obtenida) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.NivelAcademicoID, data.InstitucionAcademicaID,
		data.TituloAcademicoID, data.FechaInicio, data.FechaFinalizacion,
		data.CertificacionObtenida)
}

func updateDatosAcademicos(data DatosAcademicos) (sql.Result, error) {
	query := `
		UPDATE gestion_dato_academico 
		SET obrero_cedula = ?, nivel_academico_id = ?, institucion_academica_id = ?, 
		    titulo_academico_id = ?, fecha_inicio = ?, fecha_finalizacion = ?, 
		    certificacion_obtenida = ? 
		WHERE dato_academico_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.NivelAcademicoID, data.InstitucionAcademicaID,
		data.TituloAcademicoID, data.FechaInicio, data.FechaFinalizacion,
		data.CertificacionObtenida, data.DatosAcademicosID)
}

func deleteDatosAcademicos(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_dato_academico WHERE dato_academico_id = ?"
	return repository.ExecuteQuery(query, id)
}
