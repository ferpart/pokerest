package handlers

import (
	"github.com/ferpart/pokerest/helpers"
	"net/http"
	"strconv"
	"strings"
)

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
	pokeNames := strings.Split(r.URL.Query().Get(pokeKey), ",")
	pokeArr := helpers.GetPokemon(pokeNames)

	movesInCommon := getMovesInCommon(pokeArr, language)

	limit, _ := strconv.Atoi(r.URL.Query().Get(limitKey))

	if limit == 0 {
		limit = defaultLimit
	}

	if len(movesInCommon) > limit {
		movesInCommon = movesInCommon[:limit]
	}

	helpers.SendJSON(w, http.StatusOK, movesInCommon)

}

// getMovesInCommon will return an array of moves that appear in all of the given pokemon movesets.
// if a language is specified, they'll be translated to that language, else they'll be left in
// english.
func getMovesInCommon(pokeArr []helpers.Pokemon, lan string) []helpers.Elem {
	language := helpers.GetLanguage(lan)
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
			move := translate(language, m.Move)
			movesInCommon = append(movesInCommon, move)
		}
	}

	return movesInCommon
}

// translate takes a move, and a language paramater, with that it determines to which language the
// move should be translated to, and returns the move in the new language.
func translate(lan string, move helpers.Elem) helpers.Elem {
	if lan != "" && lan != "en" {
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
