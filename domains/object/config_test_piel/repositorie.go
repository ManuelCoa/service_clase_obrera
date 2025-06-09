package test_piel

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectTestPiel() (*sql.Rows, error) {
	query := "SELECT id_piel, nombre_piel FROM config_test_piel"
	return repository.FetchRows(query)
}

func selectTestPielID(id int) (*sql.Rows, error) {
	query := "SELECT id_piel, nombre_piel FROM config_test_piel WHERE id_piel = ?"
	return repository.FetchRows(query, id)
}

func insertTestPiel(data TestPiel) (sql.Result, error) {
	query := "INSERT INTO config_test_piel (nombre_piel) VALUES (?)"
	return repository.ExecuteQuery(query, data.NombrePiel)
}

func updateTestPiel(data TestPiel) (sql.Result, error) {
	query := "UPDATE config_test_piel SET nombre_piel = ? WHERE id_piel = ?"
	return repository.ExecuteQuery(query, data.NombrePiel, data.PielID)
}

func deleteTestPiel(id int) (sql.Result, error) {
	query := "DELETE FROM config_test_piel WHERE id_piel = ?"
	return repository.ExecuteQuery(query, id)
}
