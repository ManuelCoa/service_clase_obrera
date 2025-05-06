package centro_votacion

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectCentroVotacion() (*sql.Rows, error) {
	query := "SELECT id_centro_votacion, id_estado, id_municipio, id_parroquia, codigo_centro, nombre_centro, direccion_centro, codigo_nuevo FROM config_centro_votacion"
	return repository.FetchRows(query)
}

func selectCentroVotacionID(id int) (*sql.Rows, error) {
	query := "SELECT id_centro_votacion, id_estado,id_municipio,id_parroquia, codigo_centro, nombre_centro, direccion_centro, codigo_nuevo FROM config_centro_votacion WHERE id_centro_votacion = ?"
	return repository.FetchRows(query, id)
}

func insertCentroVotacion(data CentroVotacion) (sql.Result, error) {
	query := "INSERT INTO config_centro_votacion (id_estado, id_municipio, id_parroquia, codigo_centro, nombre_centro, direccion_centro, codigo_nuevo) VALUES (?, ?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.EstadoID, data.MunicipioID, data.ParroquiaID, data.CodigoCentro, data.NombreCentro, data.DireccionCentro, data.CodigoNuevo)
}

func updateCentroVotacion(data CentroVotacion) (sql.Result, error) {
	query := "UPDATE config_centro_votacion SET id_estado = ?, id_municipio = ?, id_parroquia = ?, codigo_centro = ?, nombre_centro = ?, direccion_centro = ?, codigo_nuevo = ? WHERE id_centro_votacion = ?"
	return repository.ExecuteQuery(query, data.EstadoID, data.MunicipioID, data.ParroquiaID, data.CodigoCentro, data.NombreCentro, data.DireccionCentro, data.CodigoNuevo, data.CentroVotacionID)
}

func deleteCentroVotacion(id int) (sql.Result, error) {
	query := "DELETE FROM config_centro_votacion WHERE id_centro_votacion = ?"
	return repository.ExecuteQuery(query, id)
}
