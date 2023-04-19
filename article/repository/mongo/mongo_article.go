package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"

	"nossobr/domain"
)

const (
	database     = "article"
	collectionDB = "article"
)

type articleRepo struct {
	mongoDB *mongo.Client
}

// NewMongoArticleRepository will create an object that represent the article.Repository interface
func NewMongoArticleRepository(mongoDB *mongo.Client) domain.ArticleRepository {
	return &articleRepo{mongoDB: mongoDB}
}
