package estructura_popular

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectEstructuraPopular() (*sql.Rows, error) {
	query := "SELECT  id_estructura_popular, nombre_estructura, descripcion FROM config_estructura_popular"
	return repository.FetchRows(query)
}

func selectEstructuraPopularID(id int) (*sql.Rows, error) {
	query := "SELECT  id_estructura_popular, nombre_estructura, descripcion FROM config_estructura_popular WHERE id_estructura_popular = ?"
	return repository.FetchRows(query, id)
}

func insertEstructuraPopular(data EstructuraPopular) (sql.Result, error) {
	query := "INSERT INTO config_estructura_popular ( nombre_estructura, descripcion) VALUES (?, ?)"
	return repository.ExecuteQuery(query,data.NombreEstructura, data.Descripcion)
}

func updateEstructuraPopular(data EstructuraPopular) (sql.Result, error) {
	query := "UPDATE config_estructura_popular SET nombre_estructura = ?, descripcion = ? WHERE id_estructura_popular = ?"
	return repository.ExecuteQuery(query,data.NombreEstructura, data.Descripcion, data.EstructuraID)
}

func deleteEstructuraPopular(id int) (sql.Result, error) {
	query := "DELETE FROM config_estructura_popular WHERE id_estructura_popular = ?"
	return repository.ExecuteQuery(query, id)
}
