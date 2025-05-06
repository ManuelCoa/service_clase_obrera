package municipios

type Municipio struct {
	MunicipioID     int    `json:"id_municipio"`
	EstadoID        int    `json:"id_estado"`
	MunicipioNombre string `json:"municipio"`
}
