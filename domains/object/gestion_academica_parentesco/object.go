package academica_parentesco

import (
	"time"
)

type AcademicaParentesco struct {
	DatosAcademicosID      int       `json:"id_datos_academicos"`
	CedulaParentesco       int       `json:"cedula_parentesco"`
	NivelAcademicoID       int       `json:"id_nivel_academico"`
	TituloAcademicoID      int       `json:"id_titulo_academico"`
	InstitucionAcademicaID int       `json:"id_institucion_academica"`
	FechaInicio           time.Time `json:"fecha_inicio"`
	FechaFin              time.Time `json:"fecha_fin"`
	Certificado           string    `json:"certificado"`
}