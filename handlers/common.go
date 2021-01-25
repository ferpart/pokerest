package handlers

import (
	"github.com/ferpart/pokerest/helpers"
	"net/http"
	"strconv"
	"strings"
)

// hasCommon is a struct that stores all the common moves in a struct a list of common moves,
// as well as the following attributes:
//		- Name:		name of the request type
//		- Language:	language in which the common moves were requested
//		- Pokemon:	pokemon whose common moves were requested
//		- Moves:	moves in common for pokemon
type hasCommon struct {
	Name     string         `json:"name"`
	Language string         `json:"language"`
	Pokemon  []string       `json:"pokemon"`
	Moves    []helpers.Elem `json:"moves"`
}

// CommonMoves returns as a json format an array of all the moves that the given pokemon have in
// common. Depending on the flags given, different pokemon can be compared, the language of the
// moves can be set, and there's also a limit of how many moves will be sent that defaults to 10
// if no custom limit is given.
func CommonMoves(w http.ResponseWriter, r *http.Request) {
	const (
		defaultLimit = 10
		defaultLang  = "en"
		pokeKey      = "pokemon"
		lanKey       = "language"
		limitKey     = "limit"
	)

	language := r.URL.Query().Get(lanKey)

	common := hasCommon{
		Name:    "Common",
		Pokemon: strings.Split(r.URL.Query().Get(pokeKey), ","),
	}

	common = getMovesInCommon(common, language)

	limit, _ := strconv.Atoi(r.URL.Query().Get(limitKey))

	if limit == 0 {
		limit = defaultLimit
	}

	if len(common.Moves) > limit {
		common.Moves = common.Moves[:limit]
	}

	helpers.SendJSON(w, http.StatusOK, common)

}

// getMovesInCommon will return a hasCommon object with the moves that appear in all of the given
//pokemon movesets. If a language is specified, they'll be translated to that language, else they'll be
//left in english.
func getMovesInCommon(commonM hasCommon, lan string) hasCommon {
	pokeArr := helpers.GetPokemon(commonM.Pokemon)
	commonM.Language = helpers.GetLanguage(lan)

	movesInCommon := []helpers.Elem{}

	smallest := 0
	for i := 0; i < len(pokeArr); i++ {
		if len(pokeArr[i].Moves) < len(pokeArr[smallest].Moves) {
			smallest = i
		}
	}
	swap(0, smallest, &pokeArr)

	for _, m := range pokeArr[0].Moves {
		common := true
		for i := 1; i < len(pokeArr); i++ {
			if moveInPokemon(m.Move.Name, pokeArr[i].Moves) != true {
				common = false
				break
			}
		}
		if common {
			move := translate(commonM.Language, m.Move)
			movesInCommon = append(movesInCommon, move)
		}
	}

	commonM.Moves = movesInCommon
	return commonM
}

// translate takes a move, and a language paramater, with that it determines to which language the
// move should be translated to, and returns the move in the new language.
func translate(lan string, move helpers.Elem) helpers.Elem {
	if lan != "en" {
		move = helpers.GetMove(move.URL, lan)
	}
	return move
}

// moveInPokemon returns true if a string given to a is found on the given list.
func moveInPokemon(a string, list []helpers.Move) bool {
	for _, b := range list {
		if b.Move.Name == a {
			return true
		}
	}
	return false
}

// swap swaps two elements of a list.
func swap(i1 int, i2 int, p *[]helpers.Pokemon) {
	tmp := (*p)[i1]
	(*p)[i1] = (*p)[i2]
	(*p)[i2] = tmp
}
