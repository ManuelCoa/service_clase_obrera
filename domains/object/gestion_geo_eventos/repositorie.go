package geo_eventos

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectGeoEventos() (*sql.Rows, error) {
    query := `SELECT id_geo_eventos, id_estado, id_municipio, 
              id_parroquia, nombre_lugar 
              FROM gestion_geo_eventos`
    return repository.FetchRows(query)
}

func selectGeoEventoID(id int) (*sql.Rows, error) {
    query := `SELECT id_geo_eventos, id_estado, id_municipio, 
              id_parroquia, nombre_lugar 
              FROM gestion_geo_eventos 
              WHERE id_geo_eventos = ?`
    return repository.FetchRows(query, id)
}

func insertGeoEvento(data GestionGeoEventos) (sql.Result, error) {
    query := `INSERT INTO gestion_geo_eventos 
              (id_estado, id_municipio, id_parroquia, nombre_lugar) 
              VALUES (?, ?, ?, ?)`
    return repository.ExecuteQuery(query, 
        data.IDEstado,
        data.IDMunicipio,
        data.IDParroquia,
        data.NombreLugar,
    )
}

func updateGeoEvento(data GestionGeoEventos) (sql.Result, error) {
    query := `UPDATE gestion_geo_eventos 
              SET id_estado = ?, id_municipio = ?, 
                  id_parroquia = ?, nombre_lugar = ? 
              WHERE id_geo_eventos = ?`
    return repository.ExecuteQuery(query,
        data.IDEstado,
        data.IDMunicipio,
        data.IDParroquia,
        data.NombreLugar,
        data.IDGeoEventos,
    )
}

func deleteGeoEvento(id int) (sql.Result, error) {
    query := "DELETE FROM gestion_geo_eventos WHERE id_geo_eventos = ?"
    return repository.ExecuteQuery(query, id)
}
