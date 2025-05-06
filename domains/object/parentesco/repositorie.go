package parentesco

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectParentesco() (*sql.Rows, error) {
	query := "SELECT  id_parentesco, nombre_parentesco FROM config_parentesco"
	return repository.FetchRows(query)
}

func selectParentescoID(id int) (*sql.Rows, error) {
	query := "SELECT  id_parentesco, nombre_parentesco FROM config_parentesco WHERE id_parentesco = ?"
	return repository.FetchRows(query, id)
}

func insertParentesco(data Parentesco) (sql.Result, error) {
	query := "INSERT INTO config_parentesco ( nombre_parentesco) VALUES (?)"
	return repository.ExecuteQuery(query,data.NombreParentesco)
}

func updateParentesco(data Parentesco) (sql.Result, error) {
	query := "UPDATE config_parentesco SET nombre_parentesco = ? WHERE id_parentesco = ?"
	return repository.ExecuteQuery(query,data.NombreParentesco, data.ParentescoID)
}

func deleteParentesco(id int) (sql.Result, error) {
	query := "DELETE FROM config_parentesco WHERE id_parentesco = ?"
	return repository.ExecuteQuery(query, id)
}
