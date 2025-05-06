package organizacion_jerarquia

type OrganizacionJerarquia struct {
	JerarquiaID     int    `json:"id_jerarquia"`
	OrganizacionID  int    `json:"id_organizacion"`
	JerarquiaNombre string `json:"nombre_jerarquia"`
}
