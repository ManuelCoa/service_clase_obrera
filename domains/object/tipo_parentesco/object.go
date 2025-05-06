package tipo_parentesco

type TipoParentesco struct {
	TipoParentescoID     int    `json:"id_tipo_parentesco"`
	NombreTipoParentesco string `json:"nombre_tipo_parentesco"`
	Descripcion          string `json:"descripcion"`
}
