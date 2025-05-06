package consejo_comunal

type ConsejoComunal struct {
	ConsejoID     int    `json:"id_consejo"`
	ComunaID      int    `json:"id_comuna"`
	NombreConsejo string `json:"nombre_consejo_comunal"`
	CodigoConsejo int    `json:"codigo_consejo"`
}