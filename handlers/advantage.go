// Package handlers has the methods used for handling the calls done to the api.
package handlers

import (
	"github.com/ferpart/pokerest/helpers"
	"net/http"
	"strings"
)

// hasAdvantage is a struct that stores the result of if the first pokemon has advantage over the
// second. It has the following fields:
// 		- Advantage:   true if first pokemon has advantage over the second.
//		- DamageTo:    is the damage multiplier that the first pokemon can do to the second.
//		- DamageFrom:  is the damage multiplier that the first pokemon can recieve from the second.
//		- Total:	   is the difference between the DamageTo and DamageFrom.
//		- Description: is the description of the shown results.
type hasAdvantage struct {
	Advantage   bool    `json:"advantage"`
	DamageTo    float64 `json:"damageTo"`
	DamageFrom  float64 `json:"damageFrom"`
	Total       float64 `json:"total"`
	Description string  `json:"description"`
}

// Advantage given two pokemon we return a json if the first has advantage over the second given two
// criteria using the pokemon's type and assuming the pokemon attacking is using a move of the same
// type as his:
// 		- Can the first deal double damage to the second?
//		- Does the first recieve half (or no) damage from the second?
// The DamageTo the second pokemon and the DamageFrom the second pokemon are substracted, if we get
// a number above 0, it means that the 1st pokemon has an advantage, if its bellow 0 means that the
// second pokemon has an advantage over the first one, finally if we get 0 we conclude that both
// pokemon are on equal ground.
func Advantage(w http.ResponseWriter, r *http.Request) {
	const (
		pokeKey = "pokemon"
	)

	pokeNames := strings.Split(r.URL.Query().Get(pokeKey), ",")
	pokeArr := helpers.GetPokemon(pokeNames)
	advantage := hasAdvantage{
		Description: "No pokemon were given",
	}
	if len(pokeArr) >= 2 {
		poke1, poke2 := pokeArr[0], pokeArr[1]
		dmgRelations := []helpers.Drelations{}

		for _, t := range poke1.Ptypes {
			relation := helpers.GetDamageRelations(t.Ptype.URL)
			dmgRelations = append(dmgRelations, relation)
		}

		advantage.DamageTo = damageTo(dmgRelations, poke2)
		advantage.DamageFrom = damageFrom(poke2, dmgRelations)
		advantage.Total = advantage.DamageTo - advantage.DamageFrom

		if advantage.Total > 0 {
			advantage.Description = poke1.Name + " has advantage over " + poke2.Name
		} else if advantage.Total < 0 {
			advantage.Description = poke1.Name + " does't have advantage over " + poke2.Name
		} else {
			advantage.Description = poke1.Name + " is in equal ground to " + poke2.Name
		}
	}

	helpers.SendJSON(w, http.StatusOK, advantage)
}

// damageTo recieves an array of damage relations as the damager, and recieves a pokemon as a
// damager. Depending if the type of the pokemon is found within the damage ralations, the damage
// value will be multiplied to see how much damage will be done from the damager to the damagee.
// This takes into consideretion the case in which a pokemon would have two types.
func damageTo(damager []helpers.Drelations, damagee helpers.Pokemon) float64 {
	damage := 1.0

	for _, t1 := range damager {
		for _, t2 := range damagee.Ptypes {
			if damage > 0 {
				if typeInDamage(t2.Ptype.Name, t1.Drelation.DoubleDto) {
					damage *= 2
					continue
				} else if typeInDamage(t2.Ptype.Name, t1.Drelation.HalfDfrom) {
					damage *= 0.5
					continue
				} else if typeInDamage(t2.Ptype.Name, t1.Drelation.NoDfrom) {
					damage *= 0
				}
			}
		}
	}

	return damage
}

// damageFrom recieves a pokemon as a damager and an array of damage relations as the damagee.
// Depending if the type of the pokemon is found on the damage relations, the damge value will be
// multiplied to see how much damage the damagee will take from the damaging pokemon.
// This takes into consideration the case in which a pokemon could have two types
func damageFrom(damager helpers.Pokemon, damagee []helpers.Drelations) float64 {
	damage := 1.0

	for _, t1 := range damagee {
		for _, t2 := range damager.Ptypes {
			if damage > 0 {
				if typeInDamage(t2.Ptype.Name, t1.Drelation.DoubleDfrom) {
					damage *= 2
					continue
				} else if typeInDamage(t2.Ptype.Name, t1.Drelation.HalfDfrom) {
					damage *= 0.5
					continue
				} else if typeInDamage(t2.Ptype.Name, t1.Drelation.NoDfrom) {
					damage *= 0
				}
			}
		}
	}

	return damage
}

// typeInDamage returns true if a string given to a is found on the given list
func typeInDamage(a string, list []helpers.Elem) bool {
	for _, b := range list {
		if b.Name == a {
			return true
		}
	}
	return false
}
