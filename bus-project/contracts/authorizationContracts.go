package contracts

type UserRegisterRequest struct {
	FirstName string `json:"first_name" xml:"first_name" validate:"required"`
	LastName  string `json:"last_name" xml:"last_name" validate:"required"`
	Email     string `json:"email" xml:"email" validate:"required,email"`
	Password  string `json:"password" xml:"password" validate:"required"`
	asdasda   string
}

type UserLoginRequest struct {
	Email    string `json:"email" xml:"email" validate:"required,email"`
	Password string `json:"password" xml:"password" validate:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token" xml:"token"`
}
