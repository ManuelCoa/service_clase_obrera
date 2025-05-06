package viviendas

type Vivienda struct {
	ConfigViviendaID int    `json:"id_config_viviendas"`
	ObreroCedula     int    `json:"cedula_obrero"`
	TipoViviendaID   int    `json:"id_tipo_vivienda"`
	TipoSueloID      int    `json:"id_tipo_suelo"`
	TipoTechoID      int    `json:"id_tipo_techo"`
	NombreVivienda   string `json:"nombre_vivienda"`
	NumeroVivienda   string `json:"numero_vivienda"`
}