package object

// CORSConfig contiene la configuración para la funcionalidad CORS.
type CORSConfig struct {
	AllowedOrigins   []string // Ejemplo: []string{"https://example.com", "https://otro.com"} o ["*"]
	AllowedMethods   []string // Ejemplo: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	AllowedHeaders   []string // Ejemplo: []string{"Content-Type", "Authorization"}
	AllowCredentials bool     // true para enviar "Access-Control-Allow-Credentials: true"
	ExposedHeaders   []string // Ejemplo: []string{"X-Custom-Header"}
	MaxAge           int      // Duración en segundos para que el navegador cachee la respuesta preflight
}
