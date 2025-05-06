package ciudades

type Ciudad struct {
	CiudadID     int    `json:"id_ciudad"`
	EstadoID     int    `json:"id_estado"`
	CiudadNombre string `json:"ciudad"`
}
