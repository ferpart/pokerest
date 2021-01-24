package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestGetPokemon(t *testing.T) {
	t.Run("Should return n pokemons if n pokemons are given", func (t *testing.T)  {
		pokemon1 := []string{
			"meganium", 
			"heracross", 
			"ribombee", 
			"skarmory", 
			"tapu-lele", 
			"spiritomb",
		}
		pokemon2 := []string{"blastoise"}


		getPokemon := GetPokemon(pokemon1)
		if assert.NotEmpty(t, getPokemon){
			assert.Equal(t, "meganium", getPokemon[0].Name)
			assert.Equal(t, len(pokemon1), len(getPokemon))
		}

		getPokemon = GetPokemon(pokemon2)
		if assert.NotEmpty(t, getPokemon){
			assert.Equal(t, 1, len(getPokemon))
		}
	})
}