package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLanguage(t *testing.T) {
	t.Run("Should return a valid name after calling api", func(t *testing.T) {
		const lan = 1

		getLanguage := GetLanguage(lan)

		if assert.NotEmpty(t, getLanguage) {
			assert.Equal(t, "ja-Hrkt", getLanguage)
		}
	})
}

func TestGetMove(t *testing.T) {
	t.Run("Should return an Elem with specified language", func(t *testing.T) {
		const (
			url = "https://pokeapi.co/api/v2/move/1"
			lan = "es"
		)
		elem := Elem{
			Name: "Destructor",
			URL:  url,
		}
		move := GetMove(url, lan)

		if assert.NotEmpty(t, move) {
			assert.Equal(t, elem, move)
		}

	})
}
