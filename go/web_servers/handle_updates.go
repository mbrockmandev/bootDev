package main

import "net/http"

func (cfg *apiConfig) handleUserUpdate(w http.ResponseWriter, r *http.Request) {
 type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
    ExpiresInSeconds int `json:"expires_in_seconds,omitempty"`
  }
	
	type response struct {
		User
	}
  
  

  respondWithJSON(w, http.StatusAccepted, response{
    User: User{
      // ID: user.ID,
      // Email: user.Email,
    },
  })
}
