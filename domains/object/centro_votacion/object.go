package centro_votacion

type CentroVotacion struct {
	CentroVotacionID int    `json:"id_centro_votacion"`
	EstadoID         int    `json:"id_estado"`
	MunicipioID      int    `json:"id_municipio"`
	ParroquiaID      int    `json:"id_parroquia"`
	CodigoCentro     int    `json:"codigo_centro"`
	NombreCentro     string `json:"nombre_centro"`
	DireccionCentro  string `json:"direccion_centro"`
	CodigoNuevo      int    `json:"codigo_nuevo"`
}