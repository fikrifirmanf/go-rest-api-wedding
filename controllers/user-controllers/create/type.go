package createUser

type CreateUser struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required,lte=8,lowercase"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
}
