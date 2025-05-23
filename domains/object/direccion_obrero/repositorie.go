package direccion_obrero

import (
	"claseobrera/domains/repository"
	"database/sql"
)

func selectDireccionObrero() (*sql.Rows, error) {
	query := `
		SELECT direccion_obrero_id, obrero_cedula,id_comuna, id_consejo, estado_id, municipio_id, 
		       parroquia_id, ciudad_id, sector_urbanismo, direccion, punto_referencia, cedula_jefe_calle, nombre_jefe_calle
		FROM gestion_direccion_obrero`
	return repository.FetchRows(query)
}

func selectDireccionObreroID(id int) (*sql.Rows, error) {
	query := `SELECT direccion_obrero_id, obrero_cedula,id_comuna, id_consejo, estado_id, municipio_id, 
		       parroquia_id, ciudad_id, sector_urbanismo, direccion, punto_referencia, cedula_jefe_calle, nombre_jefe_calle
		FROM gestion_direccion_obrero 
		WHERE direccion_obrero_id = ?`
	return repository.FetchRows(query, id)
}
//funcion para consultar por id de institucion
func selectDireccionObreroInstitucionID(id_institucion int) (*sql.Rows, error) {
	query := `SELECT 
    d.cedula_obrero,
    d.id_direccion_obrero,
    d.id_estado,
    d.id_municipio,
    d.id_parroquia,
    d.id_ciudad,
    d.id_comuna,
    d.id_consejo,
    d.sector_urbanismo,
    d.direccion,
    d.punto_referencia,
    d.cedula_jefe_calle,
    d.nombre_jefe_calle
FROM 
    gestion_datos_obrero o
JOIN 
    gestion_direccion_obrero d ON o.cedula_obrero = d.cedula_obrero
WHERE 
    o.id_institucion = ?;`
	return repository.FetchRows(query, id_institucion)
}

func insertDireccionObrero(data DireccionObrero) (sql.Result, error) {
	query := `
		INSERT INTO gestion_direccion_obrero 
		(obrero_cedula, id_comuna, id_consejo, estado_id, municipio_id, parroquia_id, ciudad_id, 
		 sector_urbanismo, direccion, punto_referencia) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.ComunaID,data.ConsejoID, data.EstadoID, data.MunicipioID, data.ParroquiaID,
		data.CiudadID, data.SectorUrbanismo, data.Direccion, data.PuntoReferencia, data.CedulaJefeCalle, data.NombreJefeCalle)
}

func updateDireccionObrero(data DireccionObrero) (sql.Result, error) {
	query := `
		UPDATE gestion_direccion_obrero 
		SET obrero_cedula = ?, id_comuna = ?, id_consejo = ?, estado_id = ?, municipio_id = ?, parroquia_id = ?, 
		    ciudad_id = ?, sector_urbanismo = ?, direccion = ?, punto_referencia = ?, cedula_jefe_calle = ?, nombre_jefe_calle = ? 
		WHERE direccion_obrero_id = ?`
	return repository.ExecuteQuery(query,
		data.ObreroCedula, data.ComunaID,data.ConsejoID, data.EstadoID, data.MunicipioID, data.ParroquiaID,
		data.CiudadID, data.SectorUrbanismo, data.Direccion, data.PuntoReferencia, data.CedulaJefeCalle, data.NombreJefeCalle)
}

func deleteDireccionObrero(id int) (sql.Result, error) {
	query := "DELETE FROM gestion_direccion_obrero WHERE direccion_obrero_id = ?"
	return repository.ExecuteQuery(query, id)
}
