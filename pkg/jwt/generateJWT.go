package jwt

import (
	"system/pkg/jwt/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID, username, email string) (string, error) {

	tk := model.Token{
		UserID: userID,
		Name:   username,
		Email:  email,
		RegisteredClaims: &jwt.RegisteredClaims{
			Issuer:    "system", // Nombre de tu aplicación
			Subject:   userID,   // ID único
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // Expiración
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	// Firma con secreto (idealmente usa una variable de entorno)
	secretKey := []byte("ola mi chula") // ¡Cambia esto por una clave segura en producción!

	return token.SignedString(secretKey)
}
