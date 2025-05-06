package organizacion_estructura

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectOrganizacionEstructura() (*sql.Rows, error) {
	query := "SELECT id_estructura, nombre_estructura, abreviatura_estructura, dominio_estructura FROM config_organizacion_estructura"
	return repository.FetchRows(query)
}

func selectOrganizacionEstructuraID(id int) (*sql.Rows, error) {
	query := "SELECT id_estructura, nombre_estructura, abreviatura_estructura, dominio_estructura FROM config_organizacion_estructura WHERE id_estructura = ?"
	return repository.FetchRows(query, id)
}

func insertOrganizacionEstructura(data OrganizacionEstructura) (sql.Result, error) {
	query := "INSERT INTO config_organizacion_estructura (nombre_estructura, abreviatura_estructura, dominio_estructura) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.EstructuraNombre, data.EstructuraAbrev, data.EstructuraDominio)
}

func updateOrganizacionEstructura(data OrganizacionEstructura) (sql.Result, error) {
	query := "UPDATE config_organizacion_estructura SET nombre_estructura = ?, abreviatura_estructura = ?, dominio_estructura = ? WHERE id_estructura = ?"
	return repository.ExecuteQuery(query, data.EstructuraNombre, data.EstructuraAbrev, data.EstructuraDominio, data.EstructuraID)
}

func deleteOrganizacionEstructura(id int) (sql.Result, error) {
	query := "DELETE FROM config_organizacion_estructura WHERE id_estructura = ?"
	return repository.ExecuteQuery(query, id)
}
