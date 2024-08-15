package models

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey []byte

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (creds *Credentials)CheckCreds() bool{
	return creds.Email == "buddha@gautam.com" && creds.Password == "Bhandina"
}

func (creds *Credentials) GetJwt() (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	
	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
	
}


