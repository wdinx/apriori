package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repository *ProductRepositoryImpl) GetAll(metadata *web.Metadata) (products *[]domain.Product, err error) {
	// Get Pagination from all products
	if err = repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Find(&products).Error; err != nil {
		return nil, err
	}
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

func (repository *ProductRepositoryImpl) GetTotalPage(limit int) (totalPage int, err error) {
	var totalData int64
	if err = repository.db.Model(&domain.Product{}).Count(&totalData).Error; err != nil {
		return 0, err
	}
	totalPage = int(totalData) / limit
	return totalPage, err
}
