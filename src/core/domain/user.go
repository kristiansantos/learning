package domain

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	Status    bool
	CreateAt  string
	UpdatedAt string
}

type Users []*User

func NewUser() *User {
	return &User{}
}
