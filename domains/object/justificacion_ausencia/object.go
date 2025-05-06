package justificacion_ausencia

import (
	"time"
)

type JustificacionAusencia struct {
	JustificacionID int       `json:"id_justificacion"`
	ObreroCedula   int       `json:"cedula_obrero"`
	FechaAusencia  time.Time `json:"fecha_ausencia"`
	Motivo         string    `json:"motivo"`
	Justificativo  bool      `json:"justificativo"`
}