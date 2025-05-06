package utils

// IsOriginAllowed verifica si el origin proporcionado se encuentra en la lista de orígenes permitidos.
// Si en AllowedOrigins se incluye "*", se permite cualquier origin.
func IsOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if allowed == "*" || allowed == origin {
			return true
		}
	}
	return false
}

// JoinStrings une un slice de cadenas en una sola cadena separada por ", ".
func JoinStrings(items []string) string {
	result := ""
	for i, s := range items {
		if i > 0 {
			result += ", "
		}
		result += s
	}
	return result
}
