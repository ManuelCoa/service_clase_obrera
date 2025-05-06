package antecedentes_medicos

import "time"

type AntecedentesMedicos struct {
	AntecedentesEnfermedadID int       `json:"id_antecedente_enfermedad"`
	CedulaObrero             int       `json:"cedula_obrero"`
	TipoEnfermedadID         int       `json:"id_tipo_enfermedad"`
	FechaDiagnostico         time.Time `json:"fecha_diagnostico"`
	DuranteTrabajo           bool      `json:"durante_trabajo"`
}
