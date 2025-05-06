package habilidades

import (
	"time"
)

type Habilidades struct {
	HabilidadID           int      `json:"id_habilidades"`
	ObreroCedula         int       `json:"cedula_obrero"`
	TipoHabilidadID      int       `json:"id_tipo_habilidad"`
	NivelExperiencia     string    `json:"nivel_experiencia"` 
	Certificado          string    `json:"certificado"`
	FechaCertificacion   time.Time `json:"fecha_certificacion"`
	InstitucionCertificadora string `json:"institucion_certificadora"`
}