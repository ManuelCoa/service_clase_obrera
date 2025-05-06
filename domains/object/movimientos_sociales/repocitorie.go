package movimientos_sociales

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectMovimientosSociales() (*sql.Rows, error) {
	query := "SELECT  id_ms, nombre_ms FROM config_movimientos_sociales"
	return repository.FetchRows(query)
}

func selectMovimientosSocialesID(id int) (*sql.Rows, error) {
	query := "SELECT  id_ms, nombre_ms FROM config_movimientos_sociales WHERE id_ms = ?"
	return repository.FetchRows(query, id)
}

func insertMovimientosSociales(data MovimientosSociales) (sql.Result, error) {
	query := "INSERT INTO config_movimientos_sociales (nombre_ms) VALUES (?)"
	return repository.ExecuteQuery(query,data.NombreMovimiento)
}

func updateMovimientosSociales(data MovimientosSociales) (sql.Result, error) {
	query := "UPDATE config_movimientos_sociales SET nombre_ms = ? WHERE id_ms = ?"
	return repository.ExecuteQuery(query, data.NombreMovimiento, data.MovimientoID)
}

func deleteMovimientosSociales(id int) (sql.Result, error) {
	query := "DELETE FROM config_movimientos_sociales WHERE id_ms = ?"
	return repository.ExecuteQuery(query, id)
}
