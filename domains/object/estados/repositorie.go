package estados

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectEstados() (*sql.Rows, error) {
	query := "SELECT id_estado, estado FROM config_geo_estados"
	return repository.FetchRows(query)
}

func selectEstadoID(id int) (*sql.Rows, error) {
	query := "SELECT id_estado, estado FROM config_geo_estados WHERE id_estado = ?"
	return repository.FetchRows(query, id)
}

func insertEstado(data Estado) (sql.Result, error) {
	query := "INSERT INTO config_geo_estados (estado) VALUES (?)"
	return repository.ExecuteQuery(query, data.EstadoNombre)
}

func updateEstado(data Estado) (sql.Result, error) {
	query := "UPDATE config_geo_estados SET estado = ? WHERE id_estado = ?"
	return repository.ExecuteQuery(query, data.EstadoNombre, data.EstadoID)
}

func deleteEstado(id int) (sql.Result, error) {
	query := "DELETE FROM config_geo_estados WHERE id_estado = ?"
	return repository.ExecuteQuery(query, id)
}
