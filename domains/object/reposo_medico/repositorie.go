package reposo_medico

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectReposos() (*sql.Rows, error) {
	query := "SELECT id_reposo_medico, cedula_obrero, fecha_inicio, fecha_fin, motivo_reposo, medico_tratante FROM config_reposo_medico"
	return repository.FetchRows(query)
}

func selectRepososID(id int) (*sql.Rows, error) {
	query := "SELECT id_reposo_medico, cedula_obrero, fecha_inicio, fecha_fin, motivo_reposo, medico_tratante FROM config_reposo_medico WHERE id_reposo_medico = ?"
	return repository.FetchRows(query, id)
}

func insertReposo(data ReposoMedico) (sql.Result, error) {
	query := "INSERT INTO config_reposo_medico (cedula_obrero, fecha_inicio, fecha_fin, motivo_reposo, medico_tratante) VALUES (?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query,
		data.CedulaObrero,
		data.FechaInicio,
		data.FechaFin,
		data.MotivoReposo,
		data.MedicoTratante,
	)
}

func updateReposo(data ReposoMedico) (sql.Result, error) {
	query := "UPDATE config_reposo_medico SET cedula_obrero = ?, fecha_inicio = ?, fecha_fin = ?, motivo_reposo = ?, medico_tratante = ? WHERE id_reposo_medico = ?"
	return repository.ExecuteQuery(query,
		data.CedulaObrero,
		data.FechaInicio,
		data.FechaFin,
		data.MotivoReposo,
		data.MedicoTratante,
		data.ReposoMedicoID,
	)
}

func deleteReposo(id int) (sql.Result, error) {
	query := "DELETE FROM config_reposo_medico WHERE id_reposo_medico = ?"
	return repository.ExecuteQuery(query, id)
}
