package salud_parentesco

type SaludParentesco struct {
	SaludParentescoID int    `json:"id_salud_parentesco"`
	ParentescoCedula  int    `json:"cedula_parentesco"`
	TipoEnfermedadID  int    `json:"id_tipo_enfermedad"`
	Tratamiento       string `json:"tratamiento"`
}