// internal/auth/middleware/auth_middleware.go
package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtener token del header
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(401, gin.H{"error": "No autorizado"})
			c.Abort()
			return
		}

		// 2. Eliminar prefijo "Bearer "
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		// 3. Validar token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			return []byte("ola mi chula"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// 4. Extraer claims y agregar al contexto
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("userID", claims["sub"])
		}

		c.Next()
	}
}
