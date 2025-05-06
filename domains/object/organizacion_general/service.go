package organizacion_general

func GetOrganizacionGeneralService() ([]OrganizacionGeneral, error) {
	return searchParsedOrganizacion()
}

func GetOrganizacionGeneralIDService(id int) ([]OrganizacionGeneral, error) {
	return searchParsedOrganizacionID(id)
}

func PostOrganizacionGeneralService(data interface{}) error {
	return insertionOrganizacion(data)
}

func PutOrganizacionGeneralService(data interface{}) error {
	return putOrganizacion(data)
}

func DeleteOrganizacionGeneralService(id int) error {
	return deletionOrganizacion(id)
}
