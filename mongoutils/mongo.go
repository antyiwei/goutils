package mongoutils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSetting struct {
	URI          string
	DatabaseName string
}

var uri string
var databaseName string

func Get() MongoSetting {
	return MongoSetting{
		URI:          uri,
		DatabaseName: databaseName,
	}
}

// ConnectToDB starts a new database connection and returns a reference to it
func ConnectToDB() (*mongo.Database, error) {
	settings := Get()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(settings.URI)
	options.SetMaxPoolSize(10)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	return client.Database(settings.DatabaseName), nil
}
