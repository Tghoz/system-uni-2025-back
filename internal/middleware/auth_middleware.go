// internal/auth/middleware/auth_middleware.go
package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Obtener token de Cookie o Authorization Header
		var tokenString string
		// Intentar primero desde cookie
		tokenString, err := c.Cookie("token")
		if err != nil {
			// Si no hay cookie, intentar desde header
			authHeader := c.GetHeader("Authorization")
			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Credenciales no proporcionadas"})
				return
			}
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		}
		// 2. Validar token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
			}
			return []byte("ola mi chula"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":  "Token inválido",
				"detail": err.Error(),
			})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expirado o inválido"})
			return
		}
		// 3. Extraer y validar claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato de claims inválido"})
			return
		}
		// Verificar claims esenciales
		requiredClaims := []string{"sub", "name", "email"}
		for _, claim := range requiredClaims {
			if _, exists := claims[claim]; !exists {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Claim faltante: %s", claim)})
				return
			}
		}
		// 4. Agregar datos al contexto
		c.Set("user", gin.H{
			"id":    claims["sub"],
			"name":  claims["name"],
			"email": claims["email"],
		})

		c.Next()
	}
}
