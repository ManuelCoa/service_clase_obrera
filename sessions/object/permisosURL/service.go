package permisosURL

func GetPermissionsByRoleID(id int) ([]PermisosURL, error) {
	return searchPermissionsByRoleID(id)
}
