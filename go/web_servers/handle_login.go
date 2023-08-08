package main

import (
	"encoding/json"
	"net/http"
	"time"

	"chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds,omitempty"`
	}
	type response struct {
		Token string `json:"token"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	user, err := cfg.DB.GetUserByEmail(params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get user")
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	token, err := auth.GenerateJwt(user, time.Second*time.Duration(params.ExpiresInSeconds))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't generate token")
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: token,
	})
}
