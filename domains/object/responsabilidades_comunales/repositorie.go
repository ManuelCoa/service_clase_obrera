package responsabilidades_comunales

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectResponsabilidades() (*sql.Rows, error) {
	query := "SELECT id_responsabilidad_comunal, nombre_responsabilidad, descripcion_responsabilidad FROM config_responsabilidades_comunales"
	return repository.FetchRows(query)
}

func selectResponsabilidadesID(id int) (*sql.Rows, error) {
	query := "SELECT id_responsabilidad_comunal, nombre_responsabilidad, descripcion_responsabilidad FROM config_responsabilidades_comunales WHERE id_responsabilidad_comunal = ?"
	return repository.FetchRows(query, id)
}

func insertResponsabilidad(data ResponsabilidadComunal) (sql.Result, error) {
	query := "INSERT INTO config_responsabilidades_comunales (nombre_responsabilidad, descripcion_responsabilidad) VALUES (?, ?)"
	return repository.ExecuteQuery(query, data.NombreResponsabilidad, data.DescripcionResponsabilidad)
}

func updateResponsabilidad(data ResponsabilidadComunal) (sql.Result, error) {
	query := "UPDATE config_responsabilidades_comunales SET nombre_responsabilidad = ?, descripcion_responsabilidad = ? WHERE id_responsabilidad_comunal = ?"
	return repository.ExecuteQuery(query, data.NombreResponsabilidad, data.DescripcionResponsabilidad, data.ResponsabilidadComunalID)
}

func deleteResponsabilidad(id int) (sql.Result, error) {
	query := "DELETE FROM config_responsabilidades_comunales WHERE id_responsabilidad_comunal = ?"
	return repository.ExecuteQuery(query, id)
}
