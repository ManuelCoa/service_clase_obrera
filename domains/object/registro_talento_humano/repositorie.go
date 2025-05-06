package registro_talento_humano

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectRegistroTalentoHumano() (*sql.Rows, error) {
	query := "SELECT registro_id, obrero_cedula, cargo_id, fecha_ingreso FROM gestion_registro_talento_humano"
	return repository.FetchRows(query)
}

func selectRegistroTalentoHumanoID(id int) (*sql.Rows, error) {
	query := "SELECT registro_id, obrero_cedula, cargo_id, fecha_ingreso FROM gestion_registro_talento_humano WHERE registro_id = ?"
	return repository.FetchRows(query, id)
}

func insertRegistroTalentoHumano(data RegistroTalentoHumano) (sql.Result, error) {
	query := "INSERT INTO gestion_registro_talento_humano (obrero_cedula, cargo_id, fecha_ingreso) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.CargoID, data.FechaIngreso)
}

func updateRegistroTalentoHumano(data RegistroTalentoHumano) (sql.Result, error) {
	query := "UPDATE gestion_registro_talento_humano SET obrero_cedula = ?, cargo_id = ?, fecha_ingreso = ? WHERE registro_id = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.CargoID, data.FechaIngreso, data.RegistroID)
}

func deleteRegistroTalentoHumano(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_registro_talento_humano WHERE registro_id = ?"
	return repository.ExecuteQuery(query, id)
}
