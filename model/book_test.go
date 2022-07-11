package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook(t *testing.T) {
	t.Run("Set book", func(t *testing.T) {
		jsonElement := `{"id": 1, "Name": "The lord of the rings", "Description": "One of the most beautiful books ever made"}`
		want := Book{Id: 1, Name: "The lord of the rings", Description: "One of the most beautiful books ever made"}

		got := Book{}
		got.SetBook(jsonElement)

		assert.Equal(t, want, got)
	})
}
