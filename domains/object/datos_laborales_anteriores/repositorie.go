package datos_laborales

import (
	"database/sql"

	"claseobrera/domains/repository"
)

func selectDatosLaborales() (*sql.Rows, error) {
	query := `
		SELECT laborales_datos_id, obrero_cedula, institucion_id, 
		       cargo_id, departamento_id, fecha_inicio, 
		       fecha_fin, motivo_egreso 
		FROM gestion_datos_laborales_anteriores`
	return repository.FetchRows(query)
}

func selectDatosLaboralesID(id int) (*sql.Rows, error) {
	query := `
		SELECT laborales_datos_id, obrero_cedula, institucion_id, 
		       cargo_id, departamento_id, fecha_inicio, 
		       fecha_fin, motivo_egreso 
		FROM gestion_datos_laborales_anteriores 
		WHERE laborales_datos_id = ?`
	return repository.FetchRows(query, id)
}

func selectDatosPorObrero(cedula int) (*sql.Rows, error) {
	query := `
		SELECT laborales_datos_id, obrero_cedula, institucion_id, 
		       cargo_id, departamento_id, fecha_inicio, 
		       fecha_fin, motivo_egreso 
		FROM gestion_datos_laborales_anteriores 
		WHERE obrero_cedula = ?`
	return repository.FetchRows(query, cedula)
}

func insertDatosLaborales(data DatosLaboralesAnteriores) (sql.Result, error) {
	query := `
		INSERT INTO gestion_datos_laborales_anteriores 
		(obrero_cedula, institucion_id, cargo_id, 
		 departamento_id, fecha_inicio, fecha_fin, motivo_egreso) 
		VALUES (?, ?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.InstitucionID, data.CargoID,
		data.DepartamentoID, data.FechaInicio, data.FechaFin,
		data.MotivoEgreso)
}

func updateDatosLaborales(data DatosLaboralesAnteriores) (sql.Result, error) {
	query := `
		UPDATE gestion_datos_laborales_anteriores 
		SET obrero_cedula = ?, institucion_id = ?, cargo_id = ?, 
		    departamento_id = ?, fecha_inicio = ?, fecha_fin = ?, 
		    motivo_egreso = ? 
		WHERE laborales_datos_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.InstitucionID, data.CargoID,
		data.DepartamentoID, data.FechaInicio, data.FechaFin,
		data.MotivoEgreso, data.LaboralesDatosID)
}

func deleteDatosLaborales(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_datos_laborales_anteriores WHERE laborales_datos_id = ?"
	return repository.ExecuteQuery(query, id)
}
