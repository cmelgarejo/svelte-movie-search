package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims JWT claims
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
