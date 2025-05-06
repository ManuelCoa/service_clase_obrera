package parroquias

type Parroquia struct {
	ParroquiaID     int    `json:"id_parroquia"`
	MunicipioID     int    `json:"id_municipio"`
	ParroquiaNombre string `json:"parroquia"`
}
