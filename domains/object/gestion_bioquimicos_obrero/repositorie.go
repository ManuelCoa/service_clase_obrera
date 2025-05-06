package bioquimicos_obrero

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectBioquimicosObreros() (*sql.Rows, error) {
	query := "SELECT id_bioquimicos_obrero, cedula_obrero, fecha_analisis, nivel_glucosa, nivel_colesterol, nivel_trigliceridos, presion_arterial FROM gestion_bioquimicos_obrero"
	return repository.FetchRows(query)
}

func selectBioquimicosObreroID(id int) (*sql.Rows, error) {
	query := "SELECT id_bioquimicos_obrero, cedula_obrero, fecha_analisis, nivel_glucosa, nivel_colesterol, nivel_trigliceridos, presion_arterial FROM gestion_bioquimicos_obrero WHERE id_bioquimicos_obrero = ?"
	return repository.FetchRows(query, id)
}

func insertBioquimicosObrero(data BioquimicosObrero) (sql.Result, error) {
	query := "INSERT INTO gestion_bioquimicos_obrero (cedula_obrero, fecha_analisis, nivel_glucosa, nivel_colesterol, nivel_trigliceridos, presion_arterial) VALUES (?, ?, ?, ?, ?, ?)"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.FechaAnalisis, data.NivelGlucosa, data.NivelColesterol, data.NivelTrigliceridos, data.PresionArterial)
}

func updateBioquimicosObrero(data BioquimicosObrero) (sql.Result, error) {
	query := "UPDATE gestion_bioquimicos_obrero SET cedula_obrero = ?, fecha_analisis = ?, nivel_glucosa = ?, nivel_colesterol = ?, nivel_trigliceridos = ?, presion_arterial = ? WHERE id_bioquimicos_obrero = ?"
	return repository.ExecuteQuery(query, data.CedulaObrero, data.FechaAnalisis, data.NivelGlucosa, data.NivelColesterol, data.NivelTrigliceridos, data.PresionArterial, data.BioquimicosObreroID)
}

func deleteBioquimicosObrero(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_bioquimicos_obrero WHERE id_bioquimicos_obrero = ?"
	return repository.ExecuteQuery(query, id)
}
