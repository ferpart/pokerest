// Package helpers stores all the helper functions that are used for parsing requests to the api.
package helpers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Drelations stores a type, and the damage relations of said type
type Drelations struct {
	Name      string    `json:"name"`
	Drelation drelation `json:"damage_relations"`
}

// drelation is a structure that stores all the damages made from, and two the type in question
type drelation struct {
	DoubleDfrom []Elem `json:"double_damage_from"`
	DoubleDto   []Elem `json:"double_damage_to"`
	HalfDfrom   []Elem `json:"half_damage_from"`
	HalfDto     []Elem `json:"half_damage_to"`
	NoDfrom     []Elem `json:"no_damage_from"`
	NoDto       []Elem `json:"no_damage_to"`
}

// GetDamageRelations gets da damage relations from the passed url and returns it in a Drelations
// struct
func GetDamageRelations(url string) Drelations {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	relation := &Drelations{}
	json.NewDecoder(response.Body).Decode(relation)

	return *relation
}
