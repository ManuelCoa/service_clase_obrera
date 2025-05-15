package transporte_publico

type TransportePublico struct {
	IDTransportePublico int    `json:"id_transporte_publico"`
	CedulaObrero        int    `json:"cedula_obrero"`
	RutaTransporte      string `json:"ruta_transporte"`
	TiempoEstimado      string `json:"tiempo_estimado"`
}