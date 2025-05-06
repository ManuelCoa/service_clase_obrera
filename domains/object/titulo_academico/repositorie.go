package titulo_academico

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectTituloAcademico() (*sql.Rows, error) {
	query := "SELECT titulo_academico_id, codigo_titulo, fecha_egreso FROM gestion_titulo_academico"
	return repository.FetchRows(query)
}

func selectTituloAcademicoID(id int) (*sql.Rows, error) {
	query := "SELECT titulo_academico_id, codigo_titulo, fecha_egreso FROM gestion_titulo_academico WHERE titulo_academico_id = ?"
	return repository.FetchRows(query, id)
}

func insertTituloAcademico(data TituloAcademico) (sql.Result, error) {
	query := "INSERT INTO gestion_titulo_academico (codigo_titulo, fecha_egreso) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.CodigoTitulo, data.FechaEgreso)
}

func updateTituloAcademico(data TituloAcademico) (sql.Result, error) {
	query := "UPDATE gestion_titulo_academico SET codigo_titulo = ?, fecha_egreso = ? WHERE titulo_academico_id = ?"
	return repository.ExecuteQuery(query, data.CodigoTitulo, data.FechaEgreso, data.TituloAcademicoID)
}

func deleteTituloAcademico(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_titulo_academico WHERE titulo_academico_id = ?"
	return repository.ExecuteQuery(query, id)
}
