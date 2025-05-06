package academica_parentesco

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectAcademicasParentesco() (*sql.Rows, error) {
	query := "SELECT id_datos_academicos, cedula_parentesco, id_nivel_academico, id_titulo_academico, id_institucion_academica, fecha_inicio, fecha_fin, certificado FROM gestion_academica_parentesco"
	return repository.FetchRows(query)
}

func selectAcademicaParentescoID(id int) (*sql.Rows, error) {
	query := "SELECT id_datos_academicos, cedula_parentesco, id_nivel_academico, id_titulo_academico, id_institucion_academica, fecha_inicio, fecha_fin, certificado FROM gestion_academica_parentesco WHERE id_datos_academicos = ?"
	return repository.FetchRows(query, id)
}

func insertAcademicaParentesco(data AcademicaParentesco) (sql.Result, error) {
	query := "INSERT INTO gestion_academica_parentesco (cedula_parentesco, id_nivel_academico, id_titulo_academico, id_institucion_academica, fecha_inicio, fecha_fin, certificado) VALUES (?, ?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.CedulaParentesco, data.NivelAcademicoID, data.TituloAcademicoID, data.InstitucionAcademicaID, data.FechaInicio, data.FechaFin, data.Certificado)
}

func updateAcademicaParentesco(data AcademicaParentesco) (sql.Result, error) {
	query := "UPDATE gestion_academica_parentesco SET cedula_parentesco = ?, id_nivel_academico = ?, id_titulo_academico = ?, id_institucion_academica = ?, fecha_inicio = ?, fecha_fin = ?, certificado = ? WHERE id_datos_academicos = ?"
	return repository.ExecuteQuery(query, data.CedulaParentesco, data.NivelAcademicoID, data.TituloAcademicoID, data.InstitucionAcademicaID, data.FechaInicio, data.FechaFin, data.Certificado, data.DatosAcademicosID)
}

func deleteAcademicaParentesco(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_academica_parentesco WHERE id_datos_academicos = ?"
	return repository.ExecuteQuery(query, id)
}
