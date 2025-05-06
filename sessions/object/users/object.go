package users

type Usuarios struct {
	UserID        int    `json:"id_user"`
	RolesID       int    `json:"id_roles"`
	InstitucionID int    `json:"id_institucion"`
	Nombre        string `json:"nombres_apellidos"`
	User          string `json:"nombre_user"`
	Password      string `json:"password_user"`
}
