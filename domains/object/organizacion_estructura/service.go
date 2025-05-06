package organizacion_estructura

func GetOrganizacionEstructuraService() ([]OrganizacionEstructura, error) {
	return searchParsedOrganizacionEstructura()
}

func GetOrganizacionEstructuraIDService(id int) ([]OrganizacionEstructura, error) {
	return searchParsedOrganizacionEstructuraID(id)
}

func PostOrganizacionEstructuraService(data interface{}) error {
	return insertionOrganizacionEstructura(data)
}

func PutOrganizacionEstructuraService(data interface{}) error {
	return putOrganizacionEstructura(data)
}

func DeleteOrganizacionEstructuraService(id int) error {
	return deletionOrganizacionEstructura(id)
}
