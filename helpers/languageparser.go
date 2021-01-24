// Package helpers stores all the helper functions that are used for parsing requests to the api.
package helpers

import (
	"fmt"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetLanguage verifies if language exists and returns the name of the language.
// In case the language does not exist, it will return an empty string.
func GetLanguage(res interface{}) string {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/language/%v", res)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	lan := &Elem{}
	json.NewDecoder(response.Body).Decode(lan)

	return lan.Name

}

// CompleteMove saves the name, and all the names in different locales of a move
type CompleteMove struct {
	Name  string     `json:"name"`
	Names []language `json:"names"`
}

// language stores in Name the name of a move in a given language, and Language stores the name, and
// url of the language being used
type language struct {
	Name     string `json:"name"`
	Language Elem   `json:"language"`
}

// GetMove gets a move with the provided url and translates it to the language given in lan
func GetMove(url string, lan string) Elem {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	move := &CompleteMove{}
	json.NewDecoder(response.Body).Decode(move)

	for _, l := range move.Names {
		if l.Language.Name == lan {
			return Elem{l.Name, url}
		}
	}
	return Elem{}
}
