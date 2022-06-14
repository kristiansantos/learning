package mongodb

import (
	"context"
	"fmt"
	"html"

	"github.com/kristiansantos/learning/src/config/environment"
	"github.com/kristiansantos/learning/src/config/initializer"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var singletonStorage *Storage = nil

type Storage struct {
	Error   error
	MongoDB *mongo.Database
}

func New(ctx context.Context) Storage {
	environment := *environment.Instance

	if singletonStorage == nil {
		mongoDB, err := connect(ctx, environment)

		singletonStorage = &Storage{
			Error:   err,
			MongoDB: mongoDB,
		}
	}

	return *singletonStorage
}

func connect(ctx context.Context, init initializer.Application) (*mongo.Database, error) {
	var mongoUri string = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s",
		init.Mongo.User, init.Mongo.Pass, init.Mongo.Host, init.Mongo.Database))

	if init.Mongo.Args != "" {
		mongoUri = html.UnescapeString(fmt.Sprintf("mongodb://%s:%s@%s/%s?%s",
			init.Mongo.User, init.Mongo.Pass, init.Mongo.Host, init.Mongo.Database, init.Mongo.Args))
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

	return client.Database(init.Mongo.Database), nil
}
