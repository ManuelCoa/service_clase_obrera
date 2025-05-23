package data_obrero

type DataObrero struct {
	CedulaObrero int `json:"cedula_obrero"`
	//CargoPublicoID int `json:"id_cargo_publico"`
	//CargoOnapreID             int       `json:"id_cargo_onapre"`
	//ResponsabilidadPoliticaID int    `json:"id_responsabilidad_politica"`
	//ResponsabilidadComunalID  int    `json:"id_responsabilidad_comunal"`
	//EstructuraPopularID       int    `json:"id_estructura_popular"`
	InstitucionID int `json:"id_institucion"`
	//ProfesionID               int    `json:"id_profesion"`
	EstadoCivil     string `json:"estado_civil"`
	PielID          string `json:"id_piel"`
	Nombres         string `json:"nombres"`
	Apellidos       string `json:"apellidos"`
	FechaNacimiento string `json:"fecha_naci"`
	Genero          string `json:"genero"`
	//TipoTransporte            string    `json:"tipo_transporte"`
	NumeroTelefono  string `json:"num_telefono"`
	Correro         string `json:"correo_electronico"`
	PuntoReferencia string `json:"punto_referencia"`
	CedulaJefeCalle int    `json:"cedula_jefe_calle"`
	NombreJefeCalle string `json:"nombre_jefe_calle"`
}
