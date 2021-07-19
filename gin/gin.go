package gin

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type GinID struct {
	PK string
	SK string
}

type Gin struct {
	ID              *GinID
	Name            string
	CountryOfOrigin string
}

type GinEntity struct {
	PK              string `dynamodbav:"PK"`
	SK              string `dynamodbav:"SK"`
	Name            string `dynamodbav:"Name"`
	CountryOfOrigin string `dynamodbav:"CountryOfOrigin"`
}

type GinRepository interface {
	// CreateGin(ctx context.Context, gin *Gin) *Gin
	GetGinByID(ctx context.Context, id GinID) (*Gin, error)
	// UpdateGin(ctx context.Context, gin *Gin) *Gin
	// DeleteGinByID(ctx context.Context, id GinID)
}

func NewGinRepository(awsConfig *aws.Config) *DynamoDBGinRepository {
	if awsConfig == nil {
		awsConfig = &aws.Config{
			Region: aws.String("eu-west-1"),
		}
	}

	s := session.Must(session.NewSession(awsConfig))

	client := dynamodb.New(s)

	r := &DynamoDBGinRepository{
		client: client,
	}

	return r
}

type DynamoDBGinRepository struct {
	client *dynamodb.DynamoDB
}

var _ GinRepository = (*DynamoDBGinRepository)(nil)

// func (g *DynamoDBGinRepository) CreateGin(ctx context.Context, gin *Gin) *Gin {
// 	return nil
// }

func (g *DynamoDBGinRepository) GetGinByID(ctx context.Context, id GinID) (*Gin, error) {
	gio, err := g.client.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("ginventory"),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(id.PK),
			},
			"SK": {
				S: aws.String(id.SK),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	var entity GinEntity

	err = dynamodbattribute.UnmarshalMap(gio.Item, &entity)

	if err != nil {
		return nil, err
	}

	ginId := &GinID{
		PK: entity.PK,
		SK: entity.SK,
	}

	gin := Gin{
		ID:              ginId,
		Name:            entity.Name,
		CountryOfOrigin: entity.CountryOfOrigin,
	}

	return &gin, nil
}

// func (g *DynamoDBGinRepository) UpdateGin(ctx context.Context, gin *Gin) *Gin {
// 	return nil
// }

// func (g *DynamoDBGinRepository) DeleteGinByID(ctx context.Context, id GinID) {
// }
