package dto

type UserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
	DeletedAt string `json:"deleted_at"`
}

type UserResponseDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
}
