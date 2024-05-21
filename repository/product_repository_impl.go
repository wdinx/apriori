package repository

import (
	"apriori-backend/model/domain"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repository *ProductRepositoryImpl) GetAll() (products *[]domain.Product, err error) {
	if err = repository.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, err
}

func (repository *ProductRepositoryImpl) GetByID(id int) (product *domain.Product, err error) {
	if err = repository.db.First(&product, "id LIKE ?", id).Error; err != nil {
		return nil, err
	}
	return product, err
}

func (repository *ProductRepositoryImpl) Create(product *domain.Product) error {
	if err := repository.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ProductRepositoryImpl) Update(product *domain.Product) error {
	if err := repository.db.Where("id LIKE ?", product.ID).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

func (repository *ProductRepositoryImpl) Delete(id int) error {
	if err := repository.db.Delete(&domain.Product{}, "id LIKE ?", id).Error; err != nil {
		return err
	}
	return nil
}
