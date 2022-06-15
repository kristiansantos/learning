package user

import (
	"context"

	"github.com/kristiansantos/learning/src/core/domain"
	"github.com/kristiansantos/learning/src/shared/database/mongodb"
	"github.com/kristiansantos/learning/src/shared/provider/logger"
	"github.com/kristiansantos/learning/src/shared/tools/namespace"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection = "users"
	Namespace  = namespace.New("infra.repositories.mongodb.user.user_repository")
)

type Repository struct {
	Context    context.Context
	Collection *mongo.Collection
	Logger     logger.ILoggerProvider
}

func UserSetup(ctx context.Context) *Repository {
	connection := mongodb.New(ctx)
	log := logger.Instance

	return &Repository{
		Context:    ctx,
		Collection: connection.MongoDB.Collection(collection),
		Logger:     log,
	}
}

func (r Repository) List(filter bson.M) (users domain.Users, err error) {
	r.Logger.Info(Namespace.Concat("List"), "")

	cursor, err := r.Collection.Find(r.Context, filter)
	if err != nil {
		return
	}

	for cursor.Next(r.Context) {
		document := &domain.User{}
		cursor.Decode(&document)
		users = append(users, document)
	}

	return
}

func (r Repository) GetBy(id string) (user domain.User, err error) {
	r.Logger.Info(Namespace.Concat("GetById"), "")

	filter := bson.M{"_id": id}

	FindError := r.Collection.FindOne(r.Context, filter).Decode(&user)
	if FindError != nil {
		return domain.User{}, FindError
	}

	return
}

func (r Repository) Create(user domain.User) (domain.User, error) {
	r.Logger.Info(Namespace.Concat("Create"), "")

	_, InsertOneError := r.Collection.InsertOne(r.Context, user)
	if InsertOneError != nil {
		return domain.User{}, InsertOneError
	}

	return r.GetBy(user.ID)
}

func (r Repository) Update(id string, user domain.User) (domain.User, error) {
	r.Logger.Info(Namespace.Concat("Update"), "")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}

	_, UpdateOneError := r.Collection.UpdateOne(r.Context, filter, update)
	if UpdateOneError != nil {
		return domain.User{}, UpdateOneError
	}

	return r.GetBy(id)
}

func (r Repository) Delete(id string) (err error) {
	r.Logger.Info(Namespace.Concat("Update"), "")

	filter := bson.M{"_id": id}

	_, deleteError := r.Collection.DeleteOne(r.Context, filter)
	if deleteError != nil {
		return deleteError
	}

	return
}
