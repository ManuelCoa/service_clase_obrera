package antecedentes_medicos

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectAntecedentesMedicos() (*sql.Rows, error) {
	query := "SELECT id_antecedente_enfermedad, cedula_obrero, id_tipo_enfermedad, fecha_diagnostico, durante_trabajo FROM gestion_antecedentes_medicos"
	return repository.FetchRows(query)
}

func selectAntecedenteMedicoID(id int) (*sql.Rows, error) {
	query := "SELECT id_antecedente_enfermedad, cedula_obrero, id_tipo_enfermedad, fecha_diagnostico, durante_trabajo FROM gestion_antecedentes_medicos WHERE id_antecedente_enfermedad = ?"
	return repository.FetchRows(query, id)
}

func insertAntecedenteMedico(data AntecedentesMedicos) (sql.Result, error) {
	query := "INSERT INTO gestion_antecedentes_medicos (cedula_obrero, id_tipo_enfermedad, fecha_diagnostico, durante_trabajo) VALUES (?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.TipoEnfermedadID, data.FechaDiagnostico, data.DuranteTrabajo)
}

func updateAntecedenteMedico(data AntecedentesMedicos) (sql.Result, error) {
	query := "UPDATE gestion_antecedentes_medicos SET cedula_obrero = ?, id_tipo_enfermedad = ?, fecha_diagnostico = ?, durante_trabajo = ? WHERE id_antecedente_enfermedad = ?"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.TipoEnfermedadID, data.FechaDiagnostico, data.DuranteTrabajo, data.AntecedentesEnfermedadID)
}

func deleteAntecedenteMedico(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_antecedentes_medicos WHERE id_antecedente_enfermedad = ?"
	return repository.ExecuteQuery(query, id)
}
