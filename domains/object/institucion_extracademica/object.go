package institucion_extracademica

type InstitucionExtracademica struct {
	InstitucionID     int    `json:"id_intitucion_extracademica"`
	NombreInstitucion string `json:"nombre_institucion"`
	Telefono          int    `json:"telefono"`
	Correo            string `json:"correo"`
}
