package transporte_publico

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectTransportePublico() (*sql.Rows, error) {
    query := "SELECT id_transporte_publico, cedula_obrero, ruta_transporte, tiempo_estimado FROM gestion_transporte_publico"
    return repository.FetchRows(query)
}

func selectTransportePublicoID(id int) (*sql.Rows, error) {
    query := "SELECT id_transporte_publico, cedula_obrero, ruta_transporte, tiempo_estimado FROM gestion_transporte_publico WHERE id_transporte_publico = ?"
    return repository.FetchRows(query, id)
}

func insertTransportePublico(data TransportePublico) (sql.Result, error) {
    query := "INSERT INTO gestion_transporte_publico (cedula_obrero, ruta_transporte, tiempo_estimado) VALUES (?, ?, ?)"
    return repository.ExecuteQuery(query, 
        data.CedulaObrero,
        data.RutaTransporte,
        data.TiempoEstimado,
    )
}

func updateTransportePublico(data TransportePublico) (sql.Result, error) {
    query := "UPDATE gestion_transporte_publico SET cedula_obrero = ?, ruta_transporte = ?, tiempo_estimado = ? WHERE id_transporte_publico = ?"
    return repository.ExecuteQuery(query,
        data.CedulaObrero,
        data.RutaTransporte,
        data.TiempoEstimado,
        data.IDTransportePublico,
    )
}

func deleteTransportePublico(id int) (sql.Result, error) {
    query := "DELETE FROM gestion_transporte_publico WHERE id_transporte_publico = ?"
    return repository.ExecuteQuery(query, id)
}
