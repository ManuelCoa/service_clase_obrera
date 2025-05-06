package constancia_trabajo

import (
	"time"
)

type ConstanciaTrabajo struct {
	ConstanciaID  int       `json:"id_constancia"`
	ObreroCedula  int       `json:"cedula_obrero"`
	InstitucionID int       `json:"id_institucion"`
	FechaEmision  time.Time `json:"fecha_emision"`
	Desempeño     string    `json:"desempeño"`
	Observaciones string    `json:"observaciones"`
}
