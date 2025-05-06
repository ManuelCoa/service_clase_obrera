package tipo_eventos

type TipoEventos struct {
	IDTipoEventos   int    `json:"id_tipo_eventos"`
	NombreTipo      string `json:"nombre_tipo"`
	DescripcionTipo string `json:"descripcion_tipo"`
}