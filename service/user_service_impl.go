package service

import (
	"apriori-backend/middleware"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"apriori-backend/util"
	"apriori-backend/util/converter"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	userRepository repository.UserRepository
	validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		userRepository: userRepository,
		validator:      validator,
	}
}

// Memvalidasi proses login user
func (service *UserServiceImpl) Login(user *web.LoginRequest) (*web.LoginResponse, error) {
	err := service.validator.Struct(user)
	if err != nil {
		return &web.LoginResponse{}, err
	}

	result, err := service.userRepository.Login(user.Username)
	if err != nil {
		return &web.LoginResponse{}, err
	}

	if err = util.CheckPassword(user.Password, result.Password); err != nil {
		return &web.LoginResponse{}, err
	}

	token, err := middleware.CreateToken(result.ID, user.Username)
	if err != nil {
		return nil, err
	}

	return converter.ToLoginResponse(result, token), nil
}

// Memvalidasi proses register user
func (service *UserServiceImpl) Register(user *web.RegisterRequest) error {
	err := service.validator.Struct(user)
	if err != nil {
		return err
	}

	user.Password, err = util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	insertUser := converter.ToUserModel(user)

	err = service.userRepository.Register(insertUser)
	if err != nil {
		return err
	}

	return nil
}
