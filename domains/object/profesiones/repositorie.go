package profesiones

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectProfesiones() (*sql.Rows, error) {
	query := "SELECT  id_profesion, nombre FROM config_profesiones"
	return repository.FetchRows(query)
}

func selectProfesionesID(id int) (*sql.Rows, error) {
	query := "SELECT  id_profesion, nombre FROM config_profesiones WHERE id_profesion = ?"
	return repository.FetchRows(query, id)
}

func insertProfesiones(data Profesiones) (sql.Result, error) {
	query := "INSERT INTO config_profesiones ( nombre) VALUES (?)"
	return repository.ExecuteQuery(query,data.NombreProfesion)
}

func updateProfesiones(data Profesiones) (sql.Result, error) {
	query := "UPDATE config_profesiones SET nombre = ? WHERE id_profesion = ?"
	return repository.ExecuteQuery(query,data.NombreProfesion, data.ProfesionID)
}

func deleteProfesiones(id int) (sql.Result, error) {
	query := "DELETE FROM config_profesiones WHERE id_profesion = ?"
	return repository.ExecuteQuery(query, id)
}
