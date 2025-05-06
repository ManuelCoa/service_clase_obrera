package tipo_eventos

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTiposEvento() (*sql.Rows, error) {
    query := `SELECT id_tipo_eventos, nombre_tipo, descripcion_tipo 
              FROM config_tipo_eventos`
    return repository.FetchRows(query)
}

func selectTipoEventoID(id int) (*sql.Rows, error) {
    query := `SELECT id_tipo_eventos, nombre_tipo, descripcion_tipo 
              FROM config_tipo_eventos 
              WHERE id_tipo_eventos = ?`
    return repository.FetchRows(query, id)
}

func insertTipoEvento(data TipoEventos) (sql.Result, error) {
    query := `INSERT INTO config_tipo_eventos 
              (nombre_tipo, descripcion_tipo) 
              VALUES (?, ?)`
    return repository.ExecuteQuery(query, 
        data.NombreTipo,
        data.DescripcionTipo,
    )
}

func updateTipoEvento(data TipoEventos) (sql.Result, error) {
    query := `UPDATE config_tipo_eventos 
              SET nombre_tipo = ?, descripcion_tipo = ? 
              WHERE id_tipo_eventos = ?`
    return repository.ExecuteQuery(query,
        data.NombreTipo,
        data.DescripcionTipo,
        data.IDTipoEventos,
    )
}

func deleteTipoEvento(id int) (sql.Result, error) {
    query := "DELETE FROM config_tipo_eventos WHERE id_tipo_eventos = ?"
    return repository.ExecuteQuery(query, id)
}
