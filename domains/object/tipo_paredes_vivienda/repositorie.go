package paredes_vivienda

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposPared() (*sql.Rows, error) {
	query := `SELECT id_tipo_pared, nombre_pared, descripcion_pared 
	          FROM config_tipo_paredes_vivienda`
	return repository.FetchRows(query)
}

func selectTipoParedID(id int) (*sql.Rows, error) {
	query := `SELECT id_tipo_pared, nombre_pared, descripcion_pared 
	          FROM config_tipo_paredes_vivienda 
	          WHERE id_tipo_pared = ?`
	return repository.FetchRows(query, id)
}

func insertTipoPared(data TipoParedesVivienda) (sql.Result, error) {
	query := `INSERT INTO config_tipo_paredes_vivienda 
	          (nombre_pared, descripcion_pared) 
	          VALUES (?, ?)`
	return repository.ExecuteQuery(query,
		data.NombrePared,
		data.DescripcionPared,
	)
}

func updateTipoPared(data TipoParedesVivienda) (sql.Result, error) {
	query := `UPDATE config_tipo_paredes_vivienda 
	          SET nombre_pared = ?, descripcion_pared = ? 
	          WHERE id_tipo_pared = ?`
	return repository.ExecuteQuery(query,
		data.NombrePared,
		data.DescripcionPared,
		data.IDTipoPared,
	)
}

func deleteTipoPared(id int) (sql.Result, error) {
	query := "DELETE FROM config_tipo_paredes_vivienda WHERE id_tipo_pared = ?"
	return repository.ExecuteQuery(query, id)
}
