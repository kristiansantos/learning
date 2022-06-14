package repository

import "github.com/kristiansantos/learning/src/core/domain"

type IUserRepository interface {
	List() (domain.Users, error)
	GetBy(id string) (domain.User, error)
	Create(id string, doc interface{}) error
	Update(id string, doc interface{}) error
	Delete(id string) error
}
