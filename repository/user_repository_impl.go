package repository

import (
	"apriori-backend/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

// Mengambil data user di database
func (repository *UserRepositoryImpl) Login(username string) (user *domain.User, err error) {
	if err = repository.DB.First(&user, "username LIKE ?", username).Error; err != nil {
		return user, err
	}
	return user, nil
}

// Menyimpan data user baru ke dalam database
func (repository *UserRepositoryImpl) Register(user *domain.User) error {
	if err := repository.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
