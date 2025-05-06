package viviendas

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectVivienda() (*sql.Rows, error) {
	query := "SELECT id_config_vivienda, obrero_cedula, id_tipo_vivienda, id_tipo_suelo, id_tipo_techo, nombre_vivienda, numero_vivienda FROM gestion_viviendas"
	return repository.FetchRows(query)
}

func selectViviendaID(id int) (*sql.Rows, error) {
	query := "SELECT id_config_vivienda, obrero_cedula, id_tipo_vivienda, id_tipo_suelo, id_tipo_techo, nombre_vivienda, numero_vivienda FROM gestion_viviendas WHERE id_config_vivienda = ?"
	return repository.FetchRows(query, id)
}

func insertVivienda(data Vivienda) (sql.Result, error) {
	query := "INSERT INTO gestion_viviendas (obrero_cedula, id_tipo_vivienda, id_tipo_suelo, id_tipo_techo, nombre_vivienda, numero_vivienda) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.TipoViviendaID, data.TipoSueloID, data.TipoTechoID, data.NombreVivienda, data.NumeroVivienda)
}

func updateVivienda(data Vivienda) (sql.Result, error) {
	query := "UPDATE gestion_viviendas SET obrero_cedula = ?, id_tipo_vivienda = ?, id_tipo_suelo = ?, id_tipo_techo = ?, nombre_vivienda = ?, numero_vivienda = ? WHERE id_config_vivienda = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.TipoViviendaID, data.TipoSueloID, data.TipoTechoID, data.NombreVivienda, data.NumeroVivienda, data.ConfigViviendaID)
}

func deleteVivienda(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_viviendas WHERE id_config_vivienda = ?"
	return repository.ExecuteQuery(query, id)
}
