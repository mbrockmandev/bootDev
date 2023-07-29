package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type DBUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type Users struct {
	Users []User `json:"users"`
}

// TODO: Continue updating this function
func (cfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := DBUser{}
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	// generate hashed password
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid email/password.")
		return
	}

	dbUser := DBUser{
		Email:    params.Email,
		Password: string(hashedPw),
	}

	user, err := cfg.DB.SaveUserToFile(dbUser, "./database.json")
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user")
		return
	}

	respondWithJSON(w, http.StatusCreated, User{
		ID:    user.ID,
		Email: user.Email,
	})
}

// TODO: Implement this function
func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server error.")
	}

	dbUser, err := cfg.DB.GetUserByEmail(params.Email, "./database.json")
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Email not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Email and password do not match")
		return
	}

	var users []DBUserToJSON
	newID := len(users) + 1

	respondWithJSON(w, http.StatusOK, User{
		ID:    newID,
		Email: dbUser.Email,
	})
}
