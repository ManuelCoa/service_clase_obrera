package institucion_academica

type InstitucionAcademica struct {
	InstitucionAcademicaID int    `json:"id_institucion_academica"`
	EstadoID               int    `json:"id_estado"`
	MunicipioID            int    `json:"id_municipio"`
	ParroquiaID            int    `json:"id_parroquia"`
	CiudadID               int    `json:"id_ciudad"`
	NombreAcademia         string `json:"nombre_academia"`
	TipoInstitucion        string `json:"tipo_institucion"`
	Telefono               string `json:"telefono"`
	Correo                 string `json:"correo"`
}