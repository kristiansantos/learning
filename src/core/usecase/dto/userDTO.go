package dto

type UserDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Status    bool   `json:"status"`
	CreateAt  string `json:"createDate"`
	UpdatedAt string `json:"updatedAt"`
}

type UsersDTO []*UserDTO
