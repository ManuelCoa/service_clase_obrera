package departamentos

type Departamentos struct {
	DepartamentoID          int    `json:"id_config_departamentos"`
	InstitucionID           int    `json:"id_institucion"`
	NombreDepartamento      string `json:"nombre_departamento"`
	DescripcionDepartamento string `json:"descripcion_departamento"`
}
