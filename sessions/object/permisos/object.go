package permisos

type Permisos struct {
	PermisoID          int    `json:"id_permisos"`
	PermisoNombre      string `json:"nombre"`
	PermisoDescripcion string `json:"descripcion"`
}
