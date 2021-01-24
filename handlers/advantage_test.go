package handlers

import(
	"github.com/stretchr/testify/assert"
	"github.com/ferpart/pokerest/helpers"
	"testing"
)

func TestDamageTo(t *testing.T) {
	t.Run("Should return a float64 value with the damage multiplier", func(t *testing.T){
		pokemon := []string{"squirtle", "rhydon"} 
		pokeArr := helpers.GetPokemon(pokemon)
		dmgRelations := []helpers.Drelations{}
		poke1, poke2 := pokeArr[0], pokeArr[1]
		for _, t := range poke1.Ptypes{
			relation := helpers.GetDamageRelations(t.Ptype.URL)
			dmgRelations = append(dmgRelations, relation)
		}
		damage := damageTo(dmgRelations, poke2)

		if assert.NotEmpty(t, pokeArr){
			assert.Equal(t, "squirtle", pokeArr[0].Name)
			assert.Equal(t, len(pokemon), len(pokeArr))
		}

		if assert.NotEmpty(t, dmgRelations[0].Name){
			assert.Equal(t, "water", dmgRelations[0].Name)
		}

		if assert.NotEmpty(t,  dmgRelations[0].Drelation){
			assert.Equal(t, "grass",  dmgRelations[0].Drelation.DoubleDfrom[0].Name)
		}

		assert.Equal(t, 4.0, damage)

	})
}
func TestDamageFrom(t *testing.T) {
	t.Run("Should return a float64 value with the damage multiplier", func(t *testing.T){
		pokemon := []string{"squirtle", "rhydon"} 
		pokeArr := helpers.GetPokemon(pokemon)
		dmgRelations := []helpers.Drelations{}
		poke1, poke2 := pokeArr[0], pokeArr[1]
		for _, t := range poke1.Ptypes{
			relation := helpers.GetDamageRelations(t.Ptype.URL)
			dmgRelations = append(dmgRelations, relation)
		}
		damage := damageFrom(poke2, dmgRelations)

		if assert.NotEmpty(t, pokeArr){
			assert.Equal(t, "rhydon", pokeArr[1].Name)
			assert.Equal(t, len(pokemon), len(pokeArr))
		}

		if assert.NotEmpty(t, dmgRelations[0].Name){
			assert.Equal(t, "water", dmgRelations[0].Name)
		}

		if assert.NotEmpty(t,  dmgRelations[0].Drelation){
			assert.Equal(t, "electric",  dmgRelations[0].Drelation.DoubleDfrom[1].Name)
		}

		assert.Equal(t, 1.0, damage)

	})
}
