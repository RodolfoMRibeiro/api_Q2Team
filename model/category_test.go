package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCategory(t *testing.T) {
	t.Run("Set Category", func(t *testing.T) {
		jsonElement := `{"id":1, "name":"Horror"}`

		want := Category{Id: 1, Name: "Horror"}

		got := Category{}
		got.SetCategory(jsonElement)

		assert.Equal(t, want, got)
	})
}
