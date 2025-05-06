package municipios

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectMunicipios() (*sql.Rows, error) {
	query := "SELECT id_municipio, id_estado, municipio FROM config_geo_municipios"
	return repository.FetchRows(query)
}

func selectMunicipioID(id int) (*sql.Rows, error) {
	query := "SELECT id_municipio, id_estado, municipio FROM config_geo_municipios WHERE id_municipio = ?"
	return repository.FetchRows(query, id)
}

func insertMunicipio(data Municipio) (sql.Result, error) {
	query := "INSERT INTO config_geo_municipios (id_estado, municipio) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.EstadoID, data.MunicipioNombre)
}

func updateMunicipio(data Municipio) (sql.Result, error) {
	query := "UPDATE config_geo_municipios SET id_estado = ?, municipio = ? WHERE id_municipio = ?"
	return repository.ExecuteQuery(query, data.EstadoID, data.MunicipioNombre, data.MunicipioID)
}

func deleteMunicipio(id int) (sql.Result, error) {
	query := "DELETE FROM config_geo_municipios WHERE id_municipio = ?"
	return repository.ExecuteQuery(query, id)
}
