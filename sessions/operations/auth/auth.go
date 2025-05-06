package auth

import (
	"errors"
	"time"

	"claseobrera/sessions/hash"
	"claseobrera/sessions/object/session"
	failsession "claseobrera/sessions/object/session_fail"
	historysession "claseobrera/sessions/object/session_history"
	"claseobrera/sessions/object/users"

	"github.com/golang-jwt/jwt"
	"github.com/mssola/user_agent"
)

var jwtSecret = []byte("clase_obrera")

func SessionLoginOperation() func(username, password, ipAddress, userAgent string) (string, error) {
	return func(username, password, ipAddress, userAgent string) (string, error) {
		// 1. Buscar usuario por nombre de usuario utilizando la nueva función
		usersFound, err := users.GetUsuariosUSERService(username)
		if err != nil {
			return "", errors.New("error al buscar el usuario")
		}

		// Verificar si se encontró algún usuario con ese nombre de usuario
		if len(usersFound) == 0 {
			// Registrar intento fallido
			failsession.PostFailedLoginService(failsession.FailedLogin{
				IPAddress:         ipAddress,
				UserAgent:         userAgent,
				AttemptedUsername: username,
			})
			return "", errors.New("usuario inválido")
		}

		// Suponemos que el primer resultado es el usuario correcto (si hay coincidencias)
		user := &usersFound[0]

		// 2. Verificar contraseña
		if !hash.CheckPassword(password, user.Password) {
			// Registrar intento fallido
			failsession.PostFailedLoginService(failsession.FailedLogin{
				IPAddress:         ipAddress,
				UserAgent:         userAgent,
				AttemptedUsername: username,
			})
			return "", errors.New("contraseña incorrecta")
		}

		// 3. Extraer información del dispositivo usando el User-Agent
		ua := user_agent.New(userAgent)
		deviceInfo := extractDeviceInfo(ua)

		// 4. Generar token JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_usuario": user.UserID,
			"exp":        time.Now().Add(time.Hour * 72).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return "", errors.New("error generando token")
		}

		// 5. Crear sesión en la base de datos con deviceInfo
		err = session.PostSessionGeneralService(session.Session{
			Token:      tokenString,
			UserID:     user.UserID,
			IPAddress:  ipAddress,
			UserAgent:  userAgent,
			DeviceInfo: deviceInfo, // Incluir información del dispositivo
			Expiration: time.Now().Add(72 * time.Hour),
		})
		if err != nil {
			return "", errors.New("error creando sesión")
		}

		// 6. Registrar inicio de sesión exitoso en el historial
		err = historysession.PostLoginHistoryService(historysession.LoginHistory{
			Token:      tokenString,
			UserID:     user.UserID,
			Success:    true,
			IPAddress:  ipAddress,
			UserAgent:  userAgent,
			DeviceInfo: deviceInfo,
			Action:     "login", // Acción de login
		})
		if err != nil {
			return "", errors.New("error registrando el historial de inicio de sesión")
		}

		return tokenString, nil
	}
}

func SessionLogoutOperation() func(token, ipAddress, userAgent string) error {
	return func(token, ipAddress, userAgent string) error {
		// 1. Eliminar la sesión activa
		err := session.DeleteSessionGeneralService(token)
		if err != nil {
			return errors.New("error al eliminar la sesión")
		}

		// 2. Extraer información del dispositivo usando el User-Agent
		ua := user_agent.New(userAgent)
		deviceInfo := extractDeviceInfo(ua)

		// 3. Obtener el UserID asociado al token (suponiendo que el token contiene esta información)
		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !parsedToken.Valid {
			return errors.New("token inválido")
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("error al extraer claims del token")
		}

		userID, ok := claims["id_usuario"].(float64) // El valor es un float64 en MapClaims
		if !ok {
			return errors.New("error al obtener el UserID del token")
		}

		// 4. Registrar el cierre de sesión en el historial
		err = historysession.PostLoginHistoryService(historysession.LoginHistory{
			Token:      token,
			UserID:     int(userID), // Convertir el float64 a int
			Success:    true,
			IPAddress:  ipAddress,
			UserAgent:  userAgent,
			DeviceInfo: deviceInfo,
			Action:     "logout", // Acción de logout
		})
		if err != nil {
			return errors.New("error registrando el historial de cierre de sesión")
		}

		return nil
	}
}

// Función auxiliar para extraer información del dispositivo
func extractDeviceInfo(ua *user_agent.UserAgent) string {
	os := ua.OS()
	browser, version := ua.Browser()
	mobile := ua.Mobile()

	deviceInfo := "OS: " + os + ", Browser: " + browser + " " + version
	if mobile {
		deviceInfo += ", Mobile: Yes"
	} else {
		deviceInfo += ", Mobile: No"
	}

	return deviceInfo
}
