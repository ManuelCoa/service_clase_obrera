package organizacion_estructura

type OrganizacionEstructura struct {
	EstructuraID      int    `json:"id_estructura"`
	EstructuraNombre  string `json:"nombre_estructura"`
	EstructuraAbrev   string `json:"abreviatura_estructura"`
	EstructuraDominio string `json:"dominio_estructura"`
}
