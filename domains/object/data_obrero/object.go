package data_obrero

import "time"

type DataObrero struct {
	CedulaObrero              int       `json:"cedula_obrero"`
	CargoPublicoID            int       `json:"id_cargo_publico"`
	CargoOnapreID             int       `json:"id_cargo_onapre"`
	ResponsabilidadPoliticaID int       `json:"responsabilidad_politica"`
	ResponsabilidadComunalID  int       `json:"responsabilidad_comunal"`
	InstitucionID             int       `json:"id_institucion"`
	ProfesionID               int       `json:"id_profesion"`
	EstadoCivil               string    `json:"estado_civil"`
	Nombres                   string    `json:"nombres"`
	Apellidos                 string    `json:"apellidos"`
	FechaNacimiento           time.Time `json:"fecha_naci"`
	Genero                    string    `json:"genero"`
	TipoTransporte            string    `json:"tipo_transporte"`
	NumeroTelefono            string    `json:"num_telefono"`
	Correro                   string    `json:"correo_electronico"`
}
