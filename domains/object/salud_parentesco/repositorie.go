package salud_parentesco

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectSaludParentesco() (*sql.Rows, error) {
	query := "SELECT salud_parentesco_id, parentesco_cedula, tipo_enfermedad_id, tratamiento FROM gestion_salud_parentesco"
	return repository.FetchRows(query)
}

func selectSaludParentescoID(id int) (*sql.Rows, error) {
	query := "SELECT salud_parentesco_id, parentesco_cedula, tipo_enfermedad_id, tratamiento FROM gestion_salud_parentesco WHERE salud_parentesco_id = ?"
	return repository.FetchRows(query, id)
}

func insertSaludParentesco(data SaludParentesco) (sql.Result, error) {
	query := "INSERT INTO gestion_salud_parentesco (parentesco_cedula, tipo_enfermedad_id, tratamiento) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.ParentescoCedula, data.TipoEnfermedadID, data.Tratamiento)
}

func updateSaludParentesco(data SaludParentesco) (sql.Result, error) {
	query := "UPDATE gestion_salud_parentesco SET parentesco_cedula = ?, tipo_enfermedad_id = ?, tratamiento = ? WHERE salud_parentesco_id = ?"
	return repository.ExecuteQuery(query, data.ParentescoCedula, data.TipoEnfermedadID, data.Tratamiento, data.SaludParentescoID)
}

func deleteSaludParentesco(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_salud_parentesco WHERE salud_parentesco_id = ?"
	return repository.ExecuteQuery(query, id)
}
