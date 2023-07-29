package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

type DB struct {
	path string
	mux  *sync.RWMutex
}

type ChirpsDbStruct struct {
	Chirps map[string]Chirp `json:"chirps"`
}

type DBUserToJSON struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	HashedPw string `json:"hashedPw"`
}

func initializeDb(fileName string) error {
	err := os.Remove(fileName)
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	fmt.Println("Database created!")
	return nil
}

// TODO:update for DBuser parameter
func (db *DB) SaveUserToFile(user DBUser, fileName string) (User, error) {
	prevContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	// check current length from file
	var users []DBUserToJSON
	if len(prevContent) > 0 {
		err = json.Unmarshal(prevContent, &users)
		if err != nil {
			log.Fatal(err)
		}
	}

	newID := len(users) + 1
	userToSave := DBUserToJSON{
		ID:       newID,
		Email:    user.Email,
		HashedPw: user.Password,
	}

	users = append(users, userToSave)

	content, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return User{ID: userToSave.ID, Email: userToSave.Email}, nil
}

func (db *DB) GetUserByEmail(email string, filePath string) (*DBUser, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var users []DBUser
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, errors.New("User not found")

}

func (db *DB) SaveJSONToFile(body string, fileName string) (Chirp, error) {
	prevContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	chirpsDbResponse := ChirpsDbStruct{}
	chirpsDbResponse.Chirps = make(map[string]Chirp)
	if len(prevContent) > 0 {
		err = json.Unmarshal(prevContent, &chirpsDbResponse)
		if err != nil {
			log.Fatal(err)
		}
	}

	newID := len(chirpsDbResponse.Chirps) + 1

	var chirp Chirp
	chirp.ID = newID
	chirp.Body = body
	chirpsDbResponse.Chirps[fmt.Sprintf("%d", newID)] = chirp

	content, err := json.Marshal(chirpsDbResponse)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return chirp, nil
}

func (db *DB) GetJSONFromFile(fileName string) ([]Chirp, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Empty JSON File")
		return nil, err
	}

	chirpsWrapper := new(ChirpsDbStruct)
	if len(content) > 0 {
		err = json.Unmarshal(content, chirpsWrapper)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
	}

	chirps := make([]Chirp, 0, len(chirpsWrapper.Chirps))
	for _, chirp := range chirpsWrapper.Chirps {
		chirps = append(chirps, chirp)
	}

	sort.Slice(chirps, func(a, b int) bool {
		return chirps[a].ID < chirps[b].ID
	})

	return chirps, nil
}
