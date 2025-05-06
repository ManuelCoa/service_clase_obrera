package organizacion_jerarquia

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectOrganizacionJerarquia() (*sql.Rows, error) {
	query := "SELECT id_jerarquia, id_organizacion, nombre_jerarquia FROM config_organizacion_jerarquia"
	return repository.FetchRows(query)
}

func selectOrganizacionJerarquiaID(id int) (*sql.Rows, error) {
	query := "SELECT id_jerarquia, id_organizacion, nombre_jerarquia FROM config_organizacion_jerarquia WHERE id_jerarquia = ?"
	return repository.FetchRows(query, id)
}

func insertOrganizacionJerarquia(data OrganizacionJerarquia) (sql.Result, error) {
	query := "INSERT INTO config_organizacion_jerarquia (id_organizacion, nombre_jerarquia) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.JerarquiaID, data.JerarquiaNombre)
}

func updateOrganizacionJerarquia(data OrganizacionJerarquia) (sql.Result, error) {
	query := "UPDATE config_organizacion_jerarquia SET id_organizacion = ?, nombre_jerarquia = ? WHERE id_jerarquia = ?"
	return repository.ExecuteQuery(query, data.OrganizacionID, data.JerarquiaNombre, data.JerarquiaID)
}

func deleteOrganizacionJerarquia(id int) (sql.Result, error) {
	query := "DELETE FROM config_organizacion_jerarquia WHERE id_jerarquia = ?"
	return repository.ExecuteQuery(query, id)
}
