package users

func GetUsuariosService() ([]Usuarios, error) {
	return searchParsedUsuarios()
}

func GetUsuariosServiceID(id int) ([]Usuarios, error) {
	return searchParsedUsuariosID(id)
}

func GetUsuariosUSERService(user string) ([]Usuarios, error) {
	return searchParsedUsuariosUSER(user)
}

func PostUsuariosService(data interface{}) error {
	return insertionUsuarios(data)
}

func PutUsuariosService(data interface{}) error {
	return putUsuarios(data)
}

func DeleteUsuariosService(id int) error {
	return deletionUsuarios(id)
}
