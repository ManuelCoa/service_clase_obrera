package middleware

// CorsOptions define la configuración de CORS a nivel de aplicación.
// Esta estructura es la que conocerá el package router, de modo que éste
// no tenga dependencia alguna con el módulo security.
type CorsOptions struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
	ExposedHeaders   []string
	MaxAge           int
}
