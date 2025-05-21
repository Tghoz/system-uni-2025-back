package handler

import (
	"net/http"
	"system/internal/auth/repo"
	"system/internal/auth/service/jwt"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

// Estructura para recibir credenciales
var credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUserHandler(authRepo repo.Auth_Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Parsear el JSON
		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(400, gin.H{"error": "Datos inválidos"})
			return
		}

		// Buscar usuario por email
		user, err := authRepo.GetUserByEmail(c.Request.Context(), credentials.Email)
		if err != nil {
			// No revelar si el usuario existe o no (por seguridad)
			c.JSON(401, gin.H{"error": "Direccion de correo no existe"})
			return
		}

		// Comparar contraseña con hash
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
			c.JSON(401, gin.H{"error": "Contraseña incorrecta"})
			return
		}

		// Generar token JWT (ejemplo usando la biblioteca "github.com/golang-jwt/jwt")
		tokenString, err := jwt.GenerateToken(user.ID.String(), user.Name, user.Email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Error al generar el token"})
			return
		}
		c.SetCookie(
			"authToken", // Nombre
			tokenString, // Valor
			3600,        // MaxAge (segundos)
			"/",         // Path (debe ser "/" para accesibilidad global)
			"localhost", // Domain (¡sin puerto!)
			false,       // Secure (true en producción)
			true,        // HttpOnly (no accesible desde JS)
		)
		c.SetSameSite(http.SameSiteLaxMode) // Usa NoneMode solo en producción con HTTPS

		c.JSON(200, gin.H{"message": "Login exitoso"})
	}
}
