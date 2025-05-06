package gestion_parentesco

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectParentesco() (*sql.Rows, error) {
	query := "SELECT parentesco_cedula, obrero_cedula, tipo_parentesco_id, nombres, apellidos, genero, fecha_nacimiento, estado_civil FROM gestion_parentesco_obrero"
	return repository.FetchRows(query)
}

func selectParentescoID(id int) (*sql.Rows, error) {
	query := "SELECT parentesco_cedula, obrero_cedula, tipo_parentesco_id, nombres, apellidos, genero, fecha_nacimiento, estado_civil FROM gestion_parentesco_obrero WHERE parentesco_cedula = ?"
	return repository.FetchRows(query, id)
}

func insertParentesco(data GestionParentesco) (sql.Result, error) {
	query := "INSERT INTO gestion_parentesco_obrero (parentesco_cedula, obrero_cedula, tipo_parentesco_id, nombres, apellidos, genero, fecha_nacimiento, estado_civil) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ParentescoCedula, data.ObreroCedula, data.TipoParentescoID, data.Nombres, data.Apellidos, data.Genero, data.FechaNacimiento, data.EstadoCivil)
}

func updateParentesco(data GestionParentesco) (sql.Result, error) {
	query := "UPDATE gestion_parentesco_obrero SET obrero_cedula = ?, tipo_parentesco_id = ?, nombres = ?, apellidos = ?, genero = ?, fecha_nacimiento = ?, estado_civil = ? WHERE parentesco_cedula = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.TipoParentescoID, data.Nombres, data.Apellidos, data.Genero, data.FechaNacimiento, data.EstadoCivil, data.ParentescoCedula)
}

func deleteParentesco(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_parentesco_obrero WHERE parentesco_cedula = ?"
	return repository.ExecuteQuery(query, id)
}
