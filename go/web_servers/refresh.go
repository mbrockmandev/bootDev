package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) handleUserRefresh(w http.ResponseWriter, r *http.Request) {

	type response struct {
		Token string `json:"token"`
	}

	bearerToken := r.Header.Get("Authorization")
	if !isValidBearerToken(bearerToken) {
		respondWithError(w, http.StatusUnauthorized, "Invalid bearer token")
		return
	}

	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")
	// tokenString := bearerToken[7:] // check the number?

	// verify token signature
	originalToken, err := jwt.ParseWithClaims(
		tokenString,
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid bearer token")
		return
	}

	// check this is a refresh token

	// check for revocations in the database for this token

	respondWithJSON(w, http.StatusOK, response{
		Token: "token here",
	})
}
