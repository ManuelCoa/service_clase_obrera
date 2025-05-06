package responsabilidades_comunales

type ResponsabilidadComunal struct {
	ResponsabilidadComunalID   int    `json:"id_resoponsabilidad_comunal"`
	NombreResponsabilidad      string `json:"nombre_respondabilidad"`
	DescripcionResponsabilidad string `json:"descripcion_responsabilidad"`
}
