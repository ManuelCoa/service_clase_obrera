package data_obrero

type DataObrero struct {
	CedulaObrero    int    `json:"cedula_obrero"`
	InstitucionID   int    `json:"id_institucion"`
	EstadoCivil     int    `json:"estado_civil"`
	PielID          int    `json:"id_piel"`
	Nombres         string `json:"nombres"`
	Apellidos       string `json:"apellidos"`
	FechaNacimiento string `json:"fecha_naci"`
	Genero          int    `json:"genero"`
	NumeroTelefono  string `json:"num_telefono"`
	Correo          string `json:"correo_electronico"`
}
