package roles

func GetRolesService() ([]Roles, error) {
	return searchParsedRoles()
}

func GetRolesServiceID(id int) ([]Roles, error) {
	return searchParsedRolesID(id)
}

func PostRolesService(data interface{}) error {
	return insertionRoles(data)
}

func PutRolesService(data interface{}) error {
	return putRoles(data)
}

func DeleteRolesService(id int) error {
	return deletionRoles(id)
}
