package converter

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
)

func ToLoginResponse(user *domain.User, token string) *web.LoginResponse {
	return &web.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}
}

func ToUserModel(user *web.RegisterRequest) *domain.User {
	return &domain.User{
		Username: user.Username,
		Password: user.Password,
	}
}
