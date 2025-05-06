package departamentos

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectDepartamentos() (*sql.Rows, error) {
	query := "SELECT id_config_departamentos, id_institucion, nombre_departamento, descripcion_departamento FROM config_departamentos"
	return repository.FetchRows(query)
}

func selectDepartamentosID(id int) (*sql.Rows, error) {
	query := "SELECT id_config_departamentos, id_institucion, nombre_departamento, descripcion_departamento FROM config_departamentos WHERE id_config_departamentos = ?"
	return repository.FetchRows(query, id)
}

func insertDepartamentos(data Departamentos) (sql.Result, error) {
	query := "INSERT INTO config_departamentos (id_institucion, nombre_departamento, descripcion_departamento) VALUES (?, ?, ?)"
	return repository.ExecuteQuery(query, data.InstitucionID, data.NombreDepartamento, data.DescripcionDepartamento)
}

func updateDepartamentos(data Departamentos) (sql.Result, error) {
	query := "UPDATE config_departamentos SET id_institucion = ?, nombre_departamento = ?, descripcion_departamento = ? WHERE id_config_departamentos = ?"
	return repository.ExecuteQuery(query, data.InstitucionID, data.NombreDepartamento, data.DescripcionDepartamento, data.DepartamentoID)
}

func deleteDepartamentos(id int) (sql.Result, error) {
	query := "DELETE FROM config_departamentos WHERE id_config_departamentos = ?"
	return repository.ExecuteQuery(query, id)
}