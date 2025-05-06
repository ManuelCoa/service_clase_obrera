package responsabilidad_politica

type ResponsabilidadPolitica struct {
	ResponsabilidadID int    `json:"id_politica"`
	Nombre            string `json:"nombre"`
	Descripcion       string `json:"descripccion"`
}