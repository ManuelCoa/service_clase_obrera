package organizacion_jerarquia

func GetOrganizacionJerarquiaService() ([]OrganizacionJerarquia, error) {
	return searchParsedOrganizacionJerarquia()
}

func GetOrganizacionJerarquiaIDService(id int) ([]OrganizacionJerarquia, error) {
	return searchParsedOrganizacionJerarquiaID(id)
}

func PostOrganizacionJerarquiaService(data interface{}) error {
	return insertionOrganizacionJerarquia(data)
}

func PutOrganizacionJerarquiaService(data interface{}) error {
	return putOrganizacionJerarquia(data)
}

func DeleteOrganizacionJerarquiaService(id int) error {
	return deletionOrganizacionJerarquia(id)
}
