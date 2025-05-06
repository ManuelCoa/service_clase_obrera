package organizacion_instituciones

type Instituciones struct {
	InstitucionID          int    `json:"id_institucion"`
	OrganizacionID         int    `json:"id_organizacion"`
	JerarquiaID            int    `json:"id_jerarquia"`
	EstructuraID           int    `json:"id_estructura"`
	InstitucionNombre      string `json:"nombre_institucion"`
	InstitucionAbreviatura string `json:"abreviatura_institucion"`
	InstitucionDominio     string `json:"dominio_institucion"`
}
