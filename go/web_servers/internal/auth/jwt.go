package auth

import (
	"chirpy/internal/database"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// takes a user, expiry, and generates a signed token
func GenerateJwt(user database.User, expiresInSeconds time.Duration) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")

	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		Subject:   strconv.Itoa(user.ID),
		Audience:  []string{},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresInSeconds)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func AuthenticateJwt(tokenString string) (*jwt.Token, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
