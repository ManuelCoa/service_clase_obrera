package interfaz

import (
	"claseobrera/sessions/operations/auth"
	"errors"
)

// Definición de las interfaces para las operaciones de autenticación
type LoginInterface func(username, password, ipAddress, userAgent string) (string, error)
type LogoutInterface func(token, ipAddress, userAgent string) error

// Función para exponer la operación de inicio de sesión
func LoginOperation() LoginInterface {
	return func(username, password, ipAddress, userAgent string) (string, error) {
		token, err := auth.SessionLoginOperation()(username, password, ipAddress, userAgent)
		if err != nil {
			return "", errors.New("error al iniciar sesión: " + err.Error())
		}
		return token, nil
	}
}

// Función para exponer la operación de cierre de sesión
func LogoutOperation() LogoutInterface {
	return func(token, ipAddress, userAgent string) error {
		err := auth.SessionLogoutOperation()(token, ipAddress, userAgent)
		if err != nil {
			return errors.New("error al cerrar sesión: " + err.Error())
		}
		return nil
	}
}
