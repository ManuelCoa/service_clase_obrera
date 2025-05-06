package gestion_horarios

import "time"

type GestionHorario struct {
	AsignacionID int       `json:"id_asignacion"`
	ObreroCedula int       `json:"cedula_obrero"`
	HorarioID    int       `json:"id_horario"`
	FechaInicio  time.Time `json:"fecha_inicio"`
	FechaFin     time.Time `json:"fecha_fin"`
}