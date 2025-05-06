package solicitud_ingreso

import (
	"time"
)

type SolicitudIngreso struct {
	SolicitudIngresoID int       `json:"id_solicitud_ingreso"`
	InstitucionID     int       `json:"id_institucion"`
	CedulaSolicitante string    `json:"cedula_solicitante"`
	NombreSolicitante string    `json:"nombre_solicitante"`
	ApellidoSolicitante string  `json:"apellido_solicitante"`
	FechaSolicitud    time.Time `json:"fecha_solicitud"`
	Estatus           string    `json:"estatus"` 
	Observaciones     string    `json:"observaciones"`
}