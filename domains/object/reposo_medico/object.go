package reposo_medico

import (
	"time"
)

type ReposoMedico struct {
	ReposoMedicoID int       `json:"id_reposo_medico"`
	CedulaObrero   int       `json:"cedula_obrero"`
	FechaInicio    time.Time `json:"fecha_inicio"`
	FechaFin       time.Time `json:"fecha_fin"`
	MotivoReposo   string    `json:"motivo_reposo"`
	MedicoTratante string    `json:"medico_tratante"`
}
