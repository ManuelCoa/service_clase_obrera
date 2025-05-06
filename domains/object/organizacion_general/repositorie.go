package organizacion_general

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectOrganizacion() (*sql.Rows, error) {
	query := "SELECT id_organizacion, nombre_organizacion FROM config_organizacion_general"
	return repository.FetchRows(query)
}

func selectOrganizacionID(id int) (*sql.Rows, error) {
	query := "SELECT id_organizacion, nombre_organizacion FROM config_organizacion_general WHERE id_organizacion = ?"
	return repository.FetchRows(query, id)
}

func insertOrganizacion(data OrganizacionGeneral) (sql.Result, error) {
	query := "INSERT INTO config_organizacion_general (nombre_organizacion) VALUES (?)"
	return repository.ExecuteQuery(query, data.OrgNombre)
}

func updateOrganizacion(data OrganizacionGeneral) (sql.Result, error) {
	query := "UPDATE config_organizacion_general SET nombre_organizacion = ? WHERE id_organizacion = ?"
	return repository.ExecuteQuery(query, data.OrgNombre, data.OrgID)
}

func deleteOrganizacion(id int) (sql.Result, error) {
	query := "DELETE FROM config_organizacion_general WHERE id_organizacion = ?"
	return repository.ExecuteQuery(query, id)
}
