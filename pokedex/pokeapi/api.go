package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const BaseURL = "https://pokeapi.co/api/v2/location/"

type PokeApiRes struct {
	Count    int     `json:"count,omitempty"`
	Next     string  `json:"next,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Results  []struct {
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"results,omitempty"`
}

// need to call to pokeapi location area endpoint to get location areas
func GetNext(nextUrl string) (*PokeApiRes, error) {
	var res *http.Response
	var err error
	if nextUrl == "" {
		res, err = http.Get(BaseURL)
	} else {
		res, err = http.Get(nextUrl)
	}
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	formattedRes := PokeApiRes{}
	err = json.Unmarshal(body, &formattedRes)

	return &formattedRes, nil
}

func GetPrevious(previousUrl string) (*PokeApiRes, error) {
	var res *http.Response
	var err error
	if previousUrl == "" {
		res, err = http.Get(BaseURL)
	} else {
		res, err = http.Get(previousUrl)
	}
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	formattedRes := PokeApiRes{}
	err = json.Unmarshal(body, &formattedRes)

	return &formattedRes, nil
}
