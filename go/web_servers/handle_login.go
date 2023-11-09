package main

import (
	"chirpy/internal/auth"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds,omitempty"`
	}
	type response struct {
		Id    int    `json:"id"`
		Email string `json:"email"`
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
		fmt.Println("params: ", params.Email, params.Password, params.ExpiresInSeconds)
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Couldn't get user: %v", err),
		)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	expiresIn := params.ExpiresInSeconds
	dayInSeconds := 86400 // 24 hours
	if expiresIn <= 0 || expiresIn > dayInSeconds {
		expiresIn = dayInSeconds
	}

	token, err := auth.GenerateJwt(user, time.Second*time.Duration(expiresIn))
	if err != nil {
		respondWithError(
			w,
			http.StatusInternalServerError,
			fmt.Sprintf("Couldn't generate token %v", err),
		)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Id:    user.ID,
		Email: user.Email,
		Token: token,
	})
}
