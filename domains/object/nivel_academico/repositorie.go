package nivel_academico

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectNivelAcademico() (*sql.Rows, error) {
	query := "SELECT  id_nivel_academico, nombre FROM config_nivel_academico"
	return repository.FetchRows(query)
}

func selectNivelAcademicoID(id int) (*sql.Rows, error) {
	query := "SELECT  id_nivel_academico, nombre FROM config_nivel_academico WHERE id_nivel_academico = ?"
	return repository.FetchRows(query, id)
}

func insertNivelAcademico(data NivelAcademico) (sql.Result, error) {
	query := "INSERT INTO config_nivel_academico ( nombre) VALUES (?)"
	return repository.ExecuteQuery(query,data.Nombre)
}

func updateNivelAcademico(data NivelAcademico) (sql.Result, error) {
	query := "UPDATE config_nivel_academico SET nombre = ? WHERE id_nivel_academico = ?"
	return repository.ExecuteQuery(query,data.Nombre, data.NivelID)
}

func deleteNivelAcademico(id int) (sql.Result, error) {
	query := "DELETE FROM config_nivel_academico WHERE id_nivel_academico = ?"
	return repository.ExecuteQuery(query, id)
}
