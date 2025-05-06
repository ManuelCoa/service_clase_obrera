package geo_eventos

type GestionGeoEventos struct {
	IDGeoEventos int    `json:"id_geo_eventos"`
	IDEstado     int    `json:"id_estado"`
	IDMunicipio  int    `json:"id_municipio"`
	IDParroquia  int    `json:"id_parroquia"`
	NombreLugar  string `json:"nombre_lugar"`
}