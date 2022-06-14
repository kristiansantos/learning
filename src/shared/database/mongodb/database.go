package mongodb

import (
	"context"
	"fmt"
	"html"

	"github.com/kristiansantos/learning/src/config/environments"
	"github.com/kristiansantos/learning/src/config/initializers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	singletonStorage *Storage = nil
)

type Storage struct {
	Error   error
	MongoDB *mongo.Database
}

func New(ctx context.Context) Storage {
	environment := *environments.Instance

	if singletonStorage == nil {
		mongoDB, err := connect(ctx, environment)

		singletonStorage = &Storage{
			Error:   err,
			MongoDB: mongoDB,
		}
	}

	return *singletonStorage
}

func connect(ctx context.Context, initializer initializers.Initializer) (*mongo.Database, error) {
	var mongoUri string = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s",
		initializer.Mongo.User, initializer.Mongo.Pass, initializer.Mongo.Host, initializer.Mongo.Database))

	if initializer.Mongo.Args != "" {
		mongoUri = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s?%s",
			initializer.Mongo.User, initializer.Mongo.Pass, initializer.Mongo.Host, initializer.Mongo.Database, initializer.Mongo.Args))
	}

	clientOptions := options.Client().ApplyURI(mongoUri)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(initializer.Mongo.Database), nil
}
