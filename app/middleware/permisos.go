package middleware

import (
	"errors"
	"strings"

	"claseobrera/app/handlers"
	sessions "claseobrera/sessions/interfaz"

	"github.com/gofiber/fiber/v2"
)

// CheckPermissions verifica que el usuario autenticado tenga permiso para acceder al recurso solicitado.
func CheckPermissions(c *fiber.Ctx) error {
	// Se espera que el middleware de autorización ya haya colocado el "user_id" en c.Locals.
	userID, ok := c.Locals("user_id").(int)

	if !ok {
		return handlers.HandleAuthError(c, errors.New("invalid user ID"))
	}

	// Obtener el usuario a través del servicio de sesiones.
	users, err := sessions.UsersIDService(userID)()
	if err != nil || len(users) == 0 {
		return handlers.HandleAuthError(c, errors.New("usuario invalido"))
	}
	user := users[0]

	// Obtener los permisos asociados al rol del usuario.
	permissions, err := sessions.PermisosURLGetServiceID(user.RolesID)()
	if err != nil {
		return handlers.HandleInternalError(c, errors.New("cannot get permissions"))
	}
	// Procesar la ruta solicitada.
	// Se elimina el prefijo "api/" y se truncan los segmentos si la ruta tiene más de un nivel.
	path := strings.TrimPrefix(strings.Trim(c.Path(), "/"), "api/")
	segments := strings.Split(path, "/")
	if len(segments) > 1 {
		path = strings.Join(segments[:len(segments)-1], "/")
	}
	// Verificar si alguno de los permisos coincide con la ruta solicitada.
	for _, perm := range permissions {
		// Se asume que cada permiso tiene un campo "URL" que contiene la ruta permitida.
		if perm.URL == path {
			return c.Next()
		}
	}

	// Si ningún permiso coincide, se retorna un error de permisos insuficientes.
	return handlers.HandlePermissionError(c, errors.New("insufficient permissions"))
}
