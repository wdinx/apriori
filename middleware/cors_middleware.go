package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func CORSMiddleware() echo.MiddlewareFunc {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Mengizinkan semua origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})
	return echo.WrapMiddleware(corsMiddleware.Handler)
}
