package handlers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMovesInCommon(t *testing.T) {
	t.Run("Should return a Elem array with all elements in common", func(t *testing.T) {
		language := "es"
		pokemon := []string{"charizard", "squirtle", "lucario", "lugia", "rayquaza", "onix"}

		common := hasCommon{
			Name:    "Common",
			Pokemon: pokemon,
		}

		common = getMovesInCommon(common, language)

		if assert.NotEmpty(t, common.Language) {
			assert.Equal(t, "es", common.Language)
		}

		if assert.NotEmpty(t, common.Moves) {
			assert.Equal(t, "Golpe Cabeza", common.Moves[0].Name)
		}
	})

}
