package bioquimicos_obrero

import (
	"time"
)

type BioquimicosObrero struct {
    BioquimicosObreroID int       `json:"id_bioquimicos_obrero"`
    CedulaObrero       int       `json:"cedula_obrero"`
    FechaAnalisis      time.Time `json:"fecha_analisis"`
    NivelGlucosa       float64   `json:"nivel_glucosa"`
    NivelColesterol    float64   `json:"nivel_colesterol"`
    NivelTrigliceridos float64   `json:"nivel_trigliceridos"`
    PresionArterial    string    `json:"presion_arterial"`
}