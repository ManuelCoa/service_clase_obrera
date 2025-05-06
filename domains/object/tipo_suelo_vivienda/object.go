package tipo_suelo_vivienda

type TipoSueloVivienda struct {
	TipoSueloViviendaID int    `json:"id_tipo_suelo"`
	NombreTipoSuelo     string `json:"nombre_tipo_suelo"`
	Descripcion         string `json:"descripcion_suelo"`
}
