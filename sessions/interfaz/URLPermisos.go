package interfaz

import (
	"claseobrera/sessions/object/permisosurl"
)

type PermisosURLInterface func() ([]permisosurl.PermisosURL, error)

func PermisosURLGetServiceID(id int) PermisosURLInterface {
	return func() ([]permisosurl.PermisosURL, error) {
		return permisosurl.GetPermissionsByRoleID(id)
	}
}
