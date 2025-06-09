package direccion_obrero

type DireccionObrero struct {
	DireccionObreroID int    `json:"id_direccion_obrero"`
	ObreroCedula      int    `json:"cedula_obrero"`
	ComunaID          int    `json:"id_comuna"`
	ConsejoID         int    `json:"id_consejo"`
	EstadoID          int    `json:"id_estado"`
	MunicipioID       int    `json:"id_municipio"`
	ParroquiaID       int    `json:"id_parroquia"`
	CiudadID          int    `json:"id_ciudad"`
	SectorUrbanismo   string `json:"sector_urbanismo"`
	Direccion         string `json:"direccion"`
	PuntoReferencia   string `json:"punto_referencia"`
	CedulaJefeCalle   int    `json:"cedula_jefe_calle"`
	NombreJefeCalle   string `json:"nombre_jefe_calle"`
}
