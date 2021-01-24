// Package helpers stores all the helper functions that are used for parsing requests to the api.
package helpers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Pokemon stores a pokemon's name, type, and moveset
type Pokemon struct {
	Name   string  `json:"name"`
	Ptypes []ptype `json:"types"`
	Moves  []Move  `json:"moves"`
}

// ptype stores a types name as well as its url
type ptype struct {
	Ptype Elem `json:"type"`
}

// Move stores an Elem containing a move's data
type Move struct {
	Move Elem `json:"move"`
}

// Elem stores a name as well as a url
type Elem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// GetPokemon given an array of two pokemon, will return an array of type Pokemon
func GetPokemon(pokeNames []string) []Pokemon {
	url := "https://pokeapi.co/api/v2/pokemon/"
	pokeArr := []Pokemon{}

	// Calling the api to get the pokemon
	for _, p := range pokeNames {
		requestURL := url + p
		response, err := http.Get(requestURL)
		if err != nil {
			log.Fatal(err)
		}
		poke := &Pokemon{}
		json.NewDecoder(response.Body).Decode(poke)
		pokeArr = append(pokeArr, *poke)
	}

	return pokeArr
}
