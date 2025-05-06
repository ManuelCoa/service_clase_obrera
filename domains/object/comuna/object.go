package comuna

type Comuna struct {
	ComunaID             int    `json:"id_comuna"`
	CentroVotacionID     int    `json:"id_centro_votacion"`
	EstadoID             int    `json:"id_estado"`
	MunicipioID          int    `json:"id_municipio"`
	ParroquiaID          int    `json:"id_parroquia"`
	NombreComuna         string `json:"nombre_comuna"`
	CodigoCircuitoComuna int    `json:"codigo_circuito_comuna"`
}