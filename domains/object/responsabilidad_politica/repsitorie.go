package responsabilidad_politica

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectResponsabilidadPolitica() (*sql.Rows, error) {
	query := "SELECT  id_politica, nombre, descripcion FROM config_responsabilidad_politica"
	return repository.FetchRows(query)
}

func selectResponsabilidadPoliticaID(id int) (*sql.Rows, error) {
	query := "SELECT  id_politica, nombre, descripcion FROM config_responsabilidad_politica WHERE id_politica = ?"
	return repository.FetchRows(query, id)
}

func insertResponsabilidadPolitica(data ResponsabilidadPolitica) (sql.Result, error) {
	query := "INSERT INTO config_responsabilidad_politica ( nombre, descripcion) VALUES (?)"
	return repository.ExecuteQuery(query,data.Nombre, data.Descripcion)
}

func updateResponsabilidadPolitica(data ResponsabilidadPolitica) (sql.Result, error) {
	query := "UPDATE config_responsabilidad_politica SET nombre = ?, descripcion = ? WHERE id_politica = ?"
	return repository.ExecuteQuery(query,data.Nombre, data.Descripcion, data.ResponsabilidadID)
}

func deleteResponsabilidadPolitica(id int) (sql.Result, error) {
	query := "DELETE FROM config_responsabilidad_politica WHERE id_politica = ?"
	return repository.ExecuteQuery(query, id)
}
