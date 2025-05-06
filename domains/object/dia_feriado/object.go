package dia_feriado

import (
	"time"
)

type DiaFeriado struct {
	FeriadoID       int       `json:"id_feriado"`
	FechaDiaFeriado time.Time `json:"fecha_dia_feriado"`
	Descripcion     string    `json:"descripcion"`
}
