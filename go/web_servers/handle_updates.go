package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
	bearerToken := r.Header.Get("Authorization")
	if !isValidBearerToken(bearerToken) {
		respondWithError(w, http.StatusUnauthorized, "Invalid bearer token")
		return
	}

	tokenString := strings.TrimPrefix(bearerToken, "Bearer ")
	// tokenString := bearerToken[7:] // check the number?

	token, err := jwt.ParseWithClaims(
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

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't get user id")
			return
		}

		type parameters struct {
			Password         string `json:"password"`
			Email            string `json:"email"`
			ExpiresInSeconds int    `json:"expires_in_seconds,omitempty"`
		}

		var params parameters
		err = json.NewDecoder(r.Body).Decode(&params)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
			return
		}

		user, err := cfg.DB.GetUser(userId)
		if err != nil {
			respondWithError(
				w,
				http.StatusInternalServerError,
				fmt.Sprintf("Couldn't get user OK? %v", err),
			)
			return
		}

		hashedPw, err := auth.HashPassword(params.Password)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't generate hashed password")
			return
		}

		user, err = cfg.DB.UpdateUser(userId, params.Email, hashedPw)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't update user")
			return
		}

		respondWithJSON(w, http.StatusOK,
			User{
				ID:    user.ID,
				Email: user.Email,
			},
		)
	} else {
		respondWithError(w, http.StatusUnauthorized, "Invalid bearer token")
	}
}

func isValidBearerToken(bearerToken string) bool {
	if len(bearerToken) < 7 || bearerToken[:7] != "Bearer " {
		return false
	}
	return true
}
