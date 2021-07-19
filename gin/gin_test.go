package gin

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"
)

func TestGetGinByID(t *testing.T) {
	id := "GIN#TANQUERAY"

	repository := NewGinRepository(&aws.Config{
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://dynamodb-local:8000"),
	})

	g, err := repository.GetGinByID(context.Background(), GinID{
		PK: id,
		SK: id,
	})

	assert.Nil(t, err)

	assert.Equal(t, "Tanqueray", g.Name)
	assert.Equal(t, "United Kingdom", g.CountryOfOrigin)
}
