package asistencia_eventos

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectAsistencias() (*sql.Rows, error) {
    query := `SELECT id_asistencia_eventos, id_evento, cedula_obrero, 
              id_institucion, participacion, observacion 
              FROM gestion_asistencia_eventos`
    return repository.FetchRows(query)
}

func selectAsistenciaID(id int) (*sql.Rows, error) {
    query := `SELECT id_asistencia_eventos, id_evento, cedula_obrero, 
              id_institucion, participacion, observacion 
              FROM gestion_asistencia_eventos 
              WHERE id_asistencia_eventos = ?`
    return repository.FetchRows(query, id)
}

func insertAsistencia(data GestionAsistenciaEventos) (sql.Result, error) {
    query := `INSERT INTO gestion_asistencia_eventos 
              (id_evento, cedula_obrero, id_institucion, participacion, observacion) 
              VALUES (?, ?, ?, ?, ?)`
    return repository.ExecuteQuery(query, 
        data.IDEvento,
        data.CedulaObrero,
        data.IDInstitucion,
        data.Participacion,
        data.Observacion,
    )
}

func updateAsistencia(data GestionAsistenciaEventos) (sql.Result, error) {
    query := `UPDATE gestion_asistencia_eventos 
              SET id_evento = ?, cedula_obrero = ?, id_institucion = ?, 
                  participacion = ?, observacion = ? 
              WHERE id_asistencia_eventos = ?`
    return repository.ExecuteQuery(query,
        data.IDEvento,
        data.CedulaObrero,
        data.IDInstitucion,
        data.Participacion,
        data.Observacion,
        data.IDAsistenciaEventos,
    )
}

func deleteAsistencia(id int) (sql.Result, error) {
    query := "DELETE FROM gestion_asistencia_eventos WHERE id_asistencia_eventos = ?"
    return repository.ExecuteQuery(query, id)
}
