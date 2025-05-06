package cargo_publico

type CargoPublico struct {
	CargoID     int    `json:"id_cargo"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
