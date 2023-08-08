package auth

import (
	"chirpy/internal/database"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(user database.User, expiresIn time.Duration) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	// claims := jwt.MapClaims{
	// 	"id":    user.ID,
	// 	"email": user.Email,
	// 	"exp":   time.Now().Add(expiresIn).Unix(),
	// }
	claims := jwt.RegisteredClaims{
		Issuer:    "chirpy",
		Subject:   string(user.ID),
		ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
		NotBefore: &jwt.NumericDate{},
		IssuedAt:  ,
		ID:        "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
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
