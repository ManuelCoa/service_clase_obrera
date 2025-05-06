package horario

import (
	"time"
)

type Horario struct {
	HorariioID    int       `json:"id_horario"`
	ConfigHorario string    `json:"config_horario"`
	HoraEntrada   time.Time `json:"hora_entrada"`
	HoraSalida    time.Time `json:"hora_salida"`
	Descripcion   string    `json:"descripcion"`
}
