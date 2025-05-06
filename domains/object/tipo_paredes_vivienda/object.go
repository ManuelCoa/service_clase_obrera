package paredes_vivienda

type TipoParedesVivienda struct {
	IDTipoPared      int    `json:"id_tipo_pared"`
	NombrePared      string `json:"nombre_pared"`
	DescripcionPared string `json:"descripcion_pared"`
}