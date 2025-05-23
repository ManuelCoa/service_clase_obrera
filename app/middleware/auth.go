package middleware

import (
	"claseobrera/app/handlers"
	"claseobrera/security/interfaz"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Define aquí tu clave secreta
var jwtSecret = []byte("clase_obrera")

// AuthRequired es el middleware de autorización que utiliza el módulo security.
func AuthRequired(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return handlers.HandleAuthError(c, errors.New("missing or malformed JWT"))
	}

	// Se invoca la función genérica de autorización del módulo security.
	claims, err := interfaz.Authorize(tokenString, jwtSecret)
	if err != nil {
		return handlers.HandleAuthError(c, err)
	}

	// Extraer el "id_usuario" de los claims y asignarlo a c.Locals para su uso posterior.
	if id, ok := claims["id_usuario"].(float64); ok {
		userID := int(id)
		c.Locals("user_id", userID)
	} else {
		fmt.Println("User ID not found in claims")
	}

	if InstitucionID, ok := claims["id_institucion"].(float64); ok {
		InstitucionID:= int(InstitucionID)
		c.Locals("id_institucion", InstitucionID)
	} else {
		fmt.Println("id_institucion not found in claims")
	}

	// Continúa con el siguiente handler.
	return c.Next()
}
