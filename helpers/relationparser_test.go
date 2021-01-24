package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDamageRelations(t *testing.T) {
	t.Run("Should return DRelations struct with given name, and non empty Drelation struct", func(t *testing.T) {
		const (
			url = "https://pokeapi.co/api/v2/type/1"
		)
		damageRelations := GetDamageRelations(url)

		if assert.NotEmpty(t, damageRelations.Name) {
			assert.Equal(t, "normal", damageRelations.Name)
		}

		if assert.NotEmpty(t, damageRelations.Drelation) {
			assert.Equal(t, "fighting", damageRelations.Drelation.DoubleDfrom[0].Name)
		}
	})
}
