package gestion_parentesco

import (
	"time"
)

type GestionParentesco struct {
	ParentescoCedula   int     `json:"cedula_parentesco"`
	ObreroCedula     int       `json:"cedula_obrero"`
	TipoParentescoID int       `json:"id_tipo_parentesco"`
	Nombres          string    `json:"nombres_parentesco"`
	Apellidos        string    `json:"apellidos_parentesco"`
	Genero           string    `json:"genero"`        
	FechaNacimiento  time.Time `json:"fecha_naci"`
	EstadoCivil      string    `json:"estado_civil"`   
}