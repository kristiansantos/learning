package domain

import "time"

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Status    bool
	CreateAt  time.Time
	UpdatedAt time.Time
}

type Users []*User

func NewUser() *User {
	return &User{}
}
