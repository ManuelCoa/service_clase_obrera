package permisos_asistencias

import (
	"time"
)

type PermisoAsistencia struct {
	PermisoAsistenciaID int      `json:"id_permisos_asistencia"`
	ObreroCedula       int       `json:"cedula_obrero"`
	FechaSolicitud     time.Time `json:"fecha_solicitud"`
	FechaInicio        time.Time `json:"fecha_inicio"`
	FechaFin           time.Time `json:"fecha_fin"`
	Motivo             string    `json:"motivo"`
	Aprobado           bool      `json:"aprobado"`
}