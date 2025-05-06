package constancia_trabajo

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectConstancias() (*sql.Rows, error) {
	query := "SELECT constancia_id, obrero_cedula, institucion_id, fecha_emision, desempeño, observaciones FROM gestion_constancia_trabajo"
	return repository.FetchRows(query)
}

func selectConstanciaID(id int) (*sql.Rows, error) {
	query := "SELECT constancia_id, obrero_cedula, institucion_id, fecha_emision, desempeño, observaciones FROM gestion_constancia_trabajo WHERE constancia_id = ?"
	return repository.FetchRows(query, id)
}

func insertConstancia(data ConstanciaTrabajo) (sql.Result, error) {
	query := "INSERT INTO gestion_constancia_trabajo (obrero_cedula, institucion_id, fecha_emision, desempeño, observaciones) VALUES (?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.InstitucionID, data.FechaEmision, data.Desempeño, data.Observaciones)
}

func updateConstancia(data ConstanciaTrabajo) (sql.Result, error) {
	query := "UPDATE gestion_constancia_trabajo SET obrero_cedula = ?, institucion_id = ?, fecha_emision = ?, desempeño = ?, observaciones = ? WHERE constancia_id = ?"
	return repository.ExecuteQuery(query, data.ObreroCedula, data.InstitucionID, data.FechaEmision, data.Desempeño, data.Observaciones, data.ConstanciaID)
}

func deleteConstancia(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_constancia_trabajo WHERE constancia_id = ?"
	return repository.ExecuteQuery(query, id)
}
