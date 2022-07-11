package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookshel(t *testing.T) {
	t.Run("Set Bookshelf", func(t *testing.T) {
		jsonElement := `{"book": 1, "category": 1}`

		want := Bookshelf{Book: 1, Category: 1}

		got := Bookshelf{}
		got.SetBookshelf(jsonElement)

		assert.Equal(t, want, got)
	})
}
