package direccion_parentesco

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectDireccionParentesco() (*sql.Rows, error) {
    query := "SELECT id_direccion_parentesco, cedula_parentesco, id_estado, id_municipio, id_parroquia, id_ciudad, sector_urbanismo, direccion, punto_referencia FROM gestion_direccion_parentesco"
    return repository.FetchRows(query)
}

func selectDireccionParentescoID(id int) (*sql.Rows, error) {
    query := "SELECT id_direccion_parentesco, cedula_parentesco, id_estado, id_municipio, id_parroquia, id_ciudad, sector_urbanismo, direccion, punto_referencia FROM gestion_direccion_parentesco WHERE id_direccion_parentesco = ?"
    return repository.FetchRows(query, id)
}

func insertDireccionParentesco(data DireccionParentesco) (sql.Result, error) {
    query := "INSERT INTO gestion_direccion_parentesco (cedula_parentesco, id_estado, id_municipio, id_parroquia, id_ciudad, sector_urbanismo, direccion, punto_referencia) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
    return repository.ExecuteQuery(query, 
        data.CedulaParentesco,
        data.EstadoID,
        data.MunicipioID,
        data.ParroquiaID,
        data.CiudadID,
        data.SectorUrbanismo,
        data.Direccion,
        data.PuntoReferencia,
    )
}

func updateDireccionParentesco(data DireccionParentesco) (sql.Result, error) {
    query := "UPDATE gestion_direccion_parentesco SET cedula_parentesco = ?, id_estado = ?, id_municipio = ?, id_parroquia = ?, id_ciudad = ?, sector_urbanismo = ?, direccion = ?, punto_referencia = ? WHERE id_direccion_parentesco = ?"
    return repository.ExecuteQuery(query,
        data.CedulaParentesco,
        data.EstadoID,
        data.MunicipioID,
        data.ParroquiaID,
        data.CiudadID,
        data.SectorUrbanismo,
        data.Direccion,
        data.PuntoReferencia,
        data.IDDireccionParentesco,
    )
}

func deleteDireccionParentesco(id int) (sql.Result, error) {
    query := "DELETE FROM gestion_direccion_parentesco WHERE id_direccion_parentesco = ?"
    return repository.ExecuteQuery(query, id)
}