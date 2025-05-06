package partido_politico

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectPartidoPolitico() (*sql.Rows, error) {
	query := "SELECT  id_pp, nombre_pp FROM config_partido_politico"
	return repository.FetchRows(query)
}

func selectPartidoPoliticoID(id int) (*sql.Rows, error) {
	query := "SELECT  id_pp, nombre_pp FROM config_partido_politico WHERE id_pp = ?"
	return repository.FetchRows(query, id)
}

func insertPartidoPolitico(data PartidoPolitico) (sql.Result, error) {
	query := "INSERT INTO config_partido_politico ( nombre_pp) VALUES (?)"
	return repository.ExecuteQuery(query,data.NombrePartido)
}

func updatePartidoPolitico(data PartidoPolitico) (sql.Result, error) {
	query := "UPDATE config_partido_politico SET nombre_pp = ? WHERE id_pp = ?"
	return repository.ExecuteQuery(query,data.NombrePartido, data.PartidoID)
}

func deletePartidoPolitico(id int) (sql.Result, error) {
	query := "DELETE FROM config_partido_politico WHERE id_pp = ?"
	return repository.ExecuteQuery(query, id)
}
