package solicitud_egreso

import (
	"time"
)

type SolicitudEgreso struct {
	SolicitudEgresoID int       `json:"id_solicitud_egreso"`
	ObreroCedula     int       `json:"cedula_obrero"`
	FechaSolicitud   time.Time `json:"fecha_solicitud"`
	Motivo           string    `json:"motivo"`
	Estatus          string    `json:"estatus"`   
	Observaciones    string    `json:"observaciones"`
}