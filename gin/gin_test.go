package gin

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGinByID(t *testing.T) {
	id := "GIN#TANQUERAY"

	repository := NewGinRepository()

	g, err := repository.GetGinByID(context.Background(), GinID{
		PK: id,
		SK: id,
	})

	assert.Nil(t, err)

	assert.Equal(t, "Tanqueray", g.Name)
	assert.Equal(t, "United Kingdom", g.CountryOfOrigin)
}
