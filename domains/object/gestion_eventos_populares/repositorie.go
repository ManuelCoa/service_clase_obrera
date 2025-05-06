package eventos_populares

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectEventos() (*sql.Rows, error) {
    query := `SELECT id_evento, id_geo_eventos, id_tipo_eventos, id_institucion, 
              nombre_evento, cuota_personas, fecha_evento, hora_evento, 
              descripcion_evento, evento_activo 
              FROM gestion_eventos_populares`
    return repository.FetchRows(query)
}

func selectEventoID(id int) (*sql.Rows, error) {
    query := `SELECT id_evento, id_geo_eventos, id_tipo_eventos, id_institucion, 
              nombre_evento, cuota_personas, fecha_evento, hora_evento, 
              descripcion_evento, evento_activo 
              FROM gestion_eventos_populares 
              WHERE id_evento = ?`
    return repository.FetchRows(query, id)
}

func insertEvento(data GestionEventosPopulares) (sql.Result, error) {
    query := `INSERT INTO gestion_eventos_populares 
              (id_geo_eventos, id_tipo_eventos, id_institucion, nombre_evento, 
               cuota_personas, fecha_evento, hora_evento, descripcion_evento, evento_activo) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
    return repository.ExecuteQuery(query, 
        data.IDGeoEventos,
        data.IDTipoEventos,
        data.IDInstitucion,
        data.NombreEvento,
        data.CuotaPersonas,
        data.FechaEvento,
        data.HoraEvento,
        data.DescripcionEvento,
        data.EventoActivo,
    )
}

func updateEvento(data GestionEventosPopulares) (sql.Result, error) {
    query := `UPDATE gestion_eventos_populares 
              SET id_geo_eventos = ?, id_tipo_eventos = ?, id_institucion = ?, 
                  nombre_evento = ?, cuota_personas = ?, fecha_evento = ?, 
                  hora_evento = ?, descripcion_evento = ?, evento_activo = ? 
              WHERE id_evento = ?`
    return repository.ExecuteQuery(query,
        data.IDGeoEventos,
        data.IDTipoEventos,
        data.IDInstitucion,
        data.NombreEvento,
        data.CuotaPersonas,
        data.FechaEvento,
        data.HoraEvento,
        data.DescripcionEvento,
        data.EventoActivo,
        data.IDEvento,
    )
}

func deleteEvento(id int) (sql.Result, error) {
    query := "DELETE FROM gestion_eventos_populares WHERE id_evento = ?"
    return repository.ExecuteQuery(query, id)
}
