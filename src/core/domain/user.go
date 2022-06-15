package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Status    bool      `bson:"status"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

type Users []*User

func NewUser() *User {
	return &User{}
}

func (u *User) Populate() {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}
