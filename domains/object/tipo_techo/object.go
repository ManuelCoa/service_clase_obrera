package tipo_techo

type TipoTecho struct {
	TipoTechoID int    `json:"id_tipo_techo"`
	NombreTecho string `json:"nombre_techo"`
	Descripcion string `json:"descripcion_techo"`
}
