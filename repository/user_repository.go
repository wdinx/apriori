package repository

import "apriori-backend/model/domain"

type UserRepository interface {
	Login(email string) (*domain.User, error)
	Register(user *domain.User) error
}
