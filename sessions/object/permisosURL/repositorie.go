package permisosurl

import (
	"database/sql"

	"claseobrera/sessions/repository"
)

func selectPermissionsByRoleID(roleID int) (*sql.Rows, error) {
	query := "SELECT cp.nombre FROM config_permisos cp JOIN config_roles_permisos crp ON cp.id_permisos = crp.id_permisos WHERE crp.id_config_rol = ?"
	return repository.FetchRows(query, roleID)
}
