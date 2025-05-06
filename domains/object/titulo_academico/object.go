package titulo_academico

import (
	"time"
)

type TituloAcademico struct {
	TituloAcademicoID int       `json:"id_titulo_academico"`
	CodigoTitulo     string    `json:"codigo_titulo"`
	FechaEgreso      time.Time `json:"fecha_egreso"`
}