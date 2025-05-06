package perfiles

func GetPerfilesService() ([]Perfiles, error) {
	return searchParsedPerfiles()
}

func PostPerfilesService(data interface{}) error {
	return insertionPerfiles(data)
}

func DeletePerfilesService(rolesID, permisosID int) error {
	return deletionPerfiles(rolesID, permisosID)
}
