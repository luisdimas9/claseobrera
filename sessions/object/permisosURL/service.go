package permisosurl

func GetPermissionsByRoleID(id int) ([]PermisosURL, error) {
	return searchPermissionsByRoleID(id)
}
