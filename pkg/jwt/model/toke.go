package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	UserID string
	Name   string
	Email  string
	*jwt.RegisteredClaims
}
