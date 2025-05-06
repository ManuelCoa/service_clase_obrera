package institucion_extracademica

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectInstituciones() (*sql.Rows, error) {
	query := "SELECT id_intitucion_extracademica, nombre_institucion, telefono, correo FROM config_institucion_extracademica"
	return repository.FetchRows(query)
}

func selectInstitucionesID(id int) (*sql.Rows, error) {
	query := "SELECT id_intitucion_extracademica, nombre_institucion, telefono, correo FROM config_institucion_extracademica WHERE id_intitucion_extracademica = ?"
	return repository.FetchRows(query, id)
}

func insertInstitucion(data InstitucionExtracademica) (sql.Result, error) {
	query := "INSERT INTO config_institucion_extracademica (nombre_institucion, telefono, correo) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.NombreInstitucion, data.Telefono, data.Correo)
}

func updateInstitucion(data InstitucionExtracademica) (sql.Result, error) {
	query := "UPDATE config_institucion_extracademica SET nombre_institucion = ?, telefono = ?, correo = ? WHERE id_intitucion_extracademica = ?"
	return repository.ExecuteQuery(query, data.NombreInstitucion, data.Telefono, data.Correo, data.InstitucionID)
}

func deleteInstitucion(id int) (sql.Result, error) {
	query := "DELETE FROM config_institucion_extracademica WHERE id_intitucion_extracademica = ?"
	return repository.ExecuteQuery(query, id)
}
