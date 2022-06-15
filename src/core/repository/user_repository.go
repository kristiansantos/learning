package repository

import (
	"github.com/kristiansantos/learning/src/core/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type IUserRepository interface {
	List(filter bson.M) (domain.Users, error)
	GetBy(id string) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Update(id string, user domain.User) (domain.User, error)
	Delete(id string) error
}
