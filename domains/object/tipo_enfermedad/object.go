package tipo_enfermedad

type TipoEnfermedad struct {
	TipoEnfermedadID      int    `json:"id_tipo_enfermedad"`
	NombreEnfermedad      string `json:"nombre_enfermedad"`
	DescripcionEnfermedad string `json:"descripcion_enfermedad"`
}
