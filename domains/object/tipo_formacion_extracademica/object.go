package tipo_formacion_extracademica

type TipoFormacionExtracademica struct {
	TipoInformacionExtracademica int    `json:"id_formacion_extracademica"`
	NombreTipo                   string `json:"nombre_tipo"`
	Descripcion                  string `json:"descripcion"`
}