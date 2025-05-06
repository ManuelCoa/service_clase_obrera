package datos_laborales

import (
	"time"
)

type DatosLaboralesAnteriores struct {
	LaboralesDatosID int       `json:"id_datos_laborales"`
	ObreroCedula     int       `json:"cedula_obrero"`
	InstitucionID    int       `json:"id_institucion"`
	CargoID          int       `json:"id_cargo"`
	DepartamentoID   int       `json:"id_departamento"`
	FechaInicio      time.Time `json:"fecha_inicio"`
	FechaFin         time.Time `json:"fecha_fin"`
	MotivoEgreso     string    `json:"motivo_egreso"`
}
