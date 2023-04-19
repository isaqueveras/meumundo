package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeout  = 30 * time.Second
	maxConnIdleTime = 3 * time.Minute
	minPoolSize     = 20
	maxPoolSize     = 300
)

var (
	dbHost = viper.GetString("database.host")
	dbPort = viper.GetString("database.port")
	dbUser = viper.GetString("database.user")
	dbPass = viper.GetString("database.pass")

	connection = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
)

// NewMongoDBConn create new MongoDB client
func NewMongoDBConn(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(connection).
			SetAuth(options.Credential{Username: dbUser, Password: dbPass}).
			SetConnectTimeout(connectTimeout).
			SetMaxConnIdleTime(maxConnIdleTime).
			SetMinPoolSize(minPoolSize).
			SetMaxPoolSize(maxPoolSize))
	if err != nil {
		return nil, err
	}

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client, nil
}
