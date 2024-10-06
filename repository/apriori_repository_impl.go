package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"errors"
	"gorm.io/gorm"
)

type AprioriRepositoryImpl struct {
	db *gorm.DB
}

func NewAprioriRepository(db *gorm.DB) AprioriRepository {
	return &AprioriRepositoryImpl{db: db}
}

func (r *AprioriRepositoryImpl) FindAll(metadata *web.Metadata) (*[]domain.AprioriResult, error) {
	//TODO implement me
	panic("implement me")
}

func (r *AprioriRepositoryImpl) Create(apriori *domain.AprioriResult) error {
	if err := r.db.Create(apriori).Error; err != nil {
		return errors.New("Error When Saving Apriori Result")
	}
	return nil
}

func (r *AprioriRepositoryImpl) GetByID(id int) (*domain.AprioriResult, error) {
	//TODO implement me
	panic("implement me")
}
