package amonestaciones

import "time"

type Amonestaciones struct {
	AmonestacionesID  int       `json:"id_amonestacion"`
	CedulaObrero      int       `json:"cedula_Obrero"`
	FechaAmonestacion time.Time `json:"fecha_amonestacion"`
	Motivo            string    `json:"motivo"`
	Sancion           string    `json:"sancion"`
}
