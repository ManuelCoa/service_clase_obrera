package registro_talento_humano

import "time"

type RegistroTalentoHumano struct {
	RegistroID    int       `json:"id_registro"`
	ObreroCedula int       `json:"cedula_obrero"`
	CargoID      int       `json:"id_cargo"`
	FechaIngreso time.Time `json:"fecha_ingreso"`
}