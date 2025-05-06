package models

import (
	"claseobrera/app/repository"
	"claseobrera/sessions/interfaz"
)

func SessionLogin(username, password, ipAddress, userAgent string) (string, error) {
	return repository.FetchSessionsService(func() (string, error) {
		// Obtener la función de login desde la interfaz
		loginFunc := interfaz.LoginOperation()

		// Invocar la función de login y obtener el token y el error
		token, err := loginFunc(username, password, ipAddress, userAgent)
		if err != nil {
			return "", err
		}

		// Retornar el token generado
		return token, nil
	})
}

func SessionLogout(token, ipAddress, userAgent string) error {
	return repository.ModifyDomainService(func() error {
		return interfaz.LogoutOperation()(token, ipAddress, userAgent)
	})
}
