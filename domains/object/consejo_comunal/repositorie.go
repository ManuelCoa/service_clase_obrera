package consejo_comunal

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectConsejoComunal() (*sql.Rows, error) {
	query := "SELECT id_consejo, id_comuna, nombre_consejo_comunal, codigo_consejo FROM config_consejo_comunal"
	return repository.FetchRows(query)
}

func selectConsejoComunalID(id int) (*sql.Rows, error) {
	query := "SELECT id_consejo, id_comuna, nombre_consejo_comunal, codigo_consejo FROM config_consejo_comunal WHERE id_consejo = ?"
	return repository.FetchRows(query, id)
}

func insertConsejoComunal(data ConsejoComunal) (sql.Result, error) {
	query := "INSERT INTO config_consejo_comunal (id_comuna, nombre_consejo_comunal, codigo_consejo) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.ComunaID, data.NombreConsejo, data.CodigoConsejo)
}

func updateConsejoComunal(data ConsejoComunal) (sql.Result, error) {
	query := "UPDATE config_consejo_comunal SET id_comuna = ?, nombre_consejo_comunal = ?, codigo_consejo = ? WHERE id_consejo = ?"
	return repository.ExecuteQuery(query, data.ComunaID, data.NombreConsejo, data.CodigoConsejo, data.ConsejoID)
}

func deleteConsejoComunal(id int) (sql.Result, error) {
	query := "DELETE FROM config_consejo_comunal WHERE id_consejo = ?"
	return repository.ExecuteQuery(query, id)
}
