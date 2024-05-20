package web

type LoginRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" form:"password" validate:"required,min=5,max=20"`
}

type RegisterRequest struct {
	Username string `json:"username" form:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" form:"password" validate:"required,min=5,max=20"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}
