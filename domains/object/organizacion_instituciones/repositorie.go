package organizacion_instituciones

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectInstituciones() (*sql.Rows, error) {
	query := "SELECT id_institucion, id_organizacion, id_jerarquia, id_estructura, nombre_institucion, abreviatura_institucion, dominio_institucion FROM config_organizacion_institucional"
	return repository.FetchRows(query)
}

func selectInstitucionesID(id int) (*sql.Rows, error) {
	query := "SELECT id_institucion, id_organizacion, id_jerarquia, id_estructura, nombre_institucion, abreviatura_institucion, dominio_institucion FROM config_organizacion_institucional WHERE id_institucion = ?"
	return repository.FetchRows(query, id)
}

func insertInstituciones(data Instituciones) (sql.Result, error) {
	query := "INSERT INTO config_organizacion_institucional (id_organizacion, id_jerarquia, id_estructura, nombre_institucion, abreviatura_institucion, dominio_institucion) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.OrganizacionID, data.JerarquiaID, data.EstructuraID, data.InstitucionNombre, data.InstitucionAbreviatura, data.InstitucionDominio)
}

func updateInstituciones(data Instituciones) (sql.Result, error) {
	query := "UPDATE config_organizacion_institucional SET id_organizacion = ?, id_jerarquia = ?, id_estructura = ?, nombre_institucion = ?, abreviatura_institucion = ?, dominio_institucion = ? WHERE id_institucion = ?"
	return repository.ExecuteQuery(query, data.OrganizacionID, data.JerarquiaID, data.EstructuraID, data.InstitucionNombre, data.InstitucionAbreviatura, data.InstitucionDominio, data.InstitucionID)
}

func deleteInstituciones(id int) (sql.Result, error) {
	query := "DELETE FROM config_organizacion_institucional WHERE id_institucion = ?"
	return repository.ExecuteQuery(query, id)
}
