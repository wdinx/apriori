package service

import "apriori-backend/model/web"

type UserService interface {
	Login(user *web.LoginRequest) (*web.LoginResponse, error)
	Register(user *web.RegisterRequest) error
}
