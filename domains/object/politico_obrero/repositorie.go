package politico_obrero

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectPoliticoObrero() (*sql.Rows, error) {
	query := "SELECT politico_obrero_id, obrero_cedula, fuerza_politica_id, pp_id, ms_id, misiones_id, registro_unoxdiez, observacion FROM gestion_politico_obrero"
	return repository.FetchRows(query)
}

func selectPoliticoObreroID(id int) (*sql.Rows, error) {
	query := "SELECT politico_obrero_id, obrero_cedula, fuerza_politica_id, pp_id, ms_id, misiones_id, registro_unoxdiez, observacion FROM gestion_politico_obrero WHERE politico_obrero_id = ?"
	return repository.FetchRows(query, id)
}

func insertPoliticoObrero(data PoliticoObrero) (sql.Result, error) {
	query := "INSERT INTO gestion_politico_obrero (obrero_cedula, fuerza_politica_id, pp_id, ms_id, misiones_id, registro_unoxdiez, observacion) VALUES (?, ?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.FuerzaPoliticaID, data.PPID, data.MSID, data.MisionesID, data.RegistroUnoxDiez, data.Observacion)
}

func updatePoliticoObrero(data PoliticoObrero) (sql.Result, error) {
	query := "UPDATE gestion_politico_obrero SET obrero_cedula = ?, fuerza_politica_id = ?, pp_id = ?, ms_id = ?, misiones_id = ?, registro_unoxdiez = ?, observacion = ? WHERE politico_obrero_id = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.FuerzaPoliticaID, data.PPID, data.MSID, data.MisionesID, data.RegistroUnoxDiez, data.Observacion, data.PoliticoObreroID)
}

func deletePoliticoObrero(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_politico_obrero WHERE politico_obrero_id = ?"
	return repository.ExecuteQuery(query, id)
}
