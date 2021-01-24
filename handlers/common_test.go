package handlers

import (
	"github.com/ferpart/pokerest/helpers"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetMovesInCommon(t *testing.T) {
	t.Run("Should return a Elem array with all elements in common", func(t *testing.T){
		language := "es"
		pokemon := []string{"charizard","squirtle","lucario","lugia","rayquaza","onix"}
		pokeArr := helpers.GetPokemon(pokemon)
		movesInCommon := getMovesInCommon(pokeArr, language)

		if assert.NotEmpty(t, pokeArr){
			assert.Equal(t, "onix", pokeArr[0].Name)
			assert.Equal(t, len(pokemon), len(pokeArr))
		}

		if assert.NotEmpty(t, movesInCommon){
			assert.Equal(t, "Golpe Cabeza", movesInCommon[0].Name)
		}
	})


}