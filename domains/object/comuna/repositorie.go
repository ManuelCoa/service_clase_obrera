package comuna

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectComuna() (*sql.Rows, error) {
	query := "SELECT id_comuna, id_centro_votacion, id_estado, id_municipio, id_parroquia, nombre_comuna, codigo_circuito_comuna FROM config_comunas"
	return repository.FetchRows(query)
}

func selectComunaID(id int) (*sql.Rows, error) {
	query := "SELECT id_comuna, id_centro_votacion, id_estado, id_municipio, id_parroquia, nombre_comuna, codigo_circuito_comuna FROM config_comunas WHERE id_comuna = ?"
	return repository.FetchRows(query, id)
}

func insertComuna(data Comuna) (sql.Result, error) {
	query := "INSERT INTO config_comunas (id_centro_votacion, id_estado, id_municipio, id_parroquia, nombre_comuna, codigo_circuito_comuna) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.CentroVotacionID, data.EstadoID, data.MunicipioID, data.ParroquiaID, data.NombreComuna, data.CodigoCircuitoComuna)
}

func updateComuna(data Comuna) (sql.Result, error) {
	query := "UPDATE config_comunas SET id_centro_votacion = ?, id_estado = ?, id_municipio = ?, id_parroquia = ?, nombre_comuna = ?, codigo_circuito_comuna = ? WHERE id_comuna = ?"
	return repository.ExecuteQuery(query, data.CentroVotacionID, data.EstadoID, data.MunicipioID, data.ParroquiaID, data.NombreComuna, data.CodigoCircuitoComuna, data.ComunaID)
}

func deleteComuna(id int) (sql.Result, error) {
	query := "DELETE FROM config_comunas WHERE id_comuna = ?"
	return repository.ExecuteQuery(query, id)
}
