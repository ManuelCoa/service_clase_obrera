package tipo_vivienda

type TipoVivienda struct {
	TipoViviendaID     int    `json:"id_tipo_vivienda"`
	NombreTipoVivienda string `json:"nombre_tipo_vivienda"`
	Descripcion        string `json:"descripcion_tipo"`
}
