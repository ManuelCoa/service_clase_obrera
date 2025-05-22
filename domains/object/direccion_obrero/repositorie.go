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
	query := `
		SELECT direccion_obrero_id, obrero_cedula,id_comuna, id_consejo, estado_id, municipio_id, 
		       parroquia_id, ciudad_id, sector_urbanismo, direccion, punto_referencia, cedula_jefe_calle, nombre_jefe_calle
		FROM gestion_direccion_obrero 
		WHERE direccion_obrero_id = ?`
	return repository.FetchRows(query, id)
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
