package mongo

import (
	"nossobr/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database     = "author"
	collectionDB = "author"
)

type authorRepo struct {
	mongoDB *mongo.Client
}

// NewMongoAuthorRepository will create an implementation of author.Repository
func NewMongoAuthorRepository(mongoDB *mongo.Client) domain.AuthorRepository {
	return &authorRepo{mongoDB: mongoDB}
}
