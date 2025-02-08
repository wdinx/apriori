package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

// Mengambil semua data product dari database
func (repository *ProductRepositoryImpl) GetAll(metadata *web.Metadata) (products *[]domain.Product, err error) {
	// Get Pagination from all products
	fmt.Println(metadata.Limit)
	if err = repository.db.Limit(metadata.Limit).Offset(metadata.Offset()).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, err
}

// Mengambil data berdasarkan ID nya dari database
func (repository *ProductRepositoryImpl) GetByID(id int) (product *domain.Product, err error) {
	if err = repository.db.First(&product, "id LIKE ?", id).Error; err != nil {
		return nil, err
	}
	return product, err
}

// Menyimpan data produk ke dalam database
func (repository *ProductRepositoryImpl) Create(product *domain.Product) error {
	if err := repository.db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

// Melakukan update data product di database
func (repository *ProductRepositoryImpl) Update(product *domain.Product) error {
	if err := repository.db.Where("id LIKE ?", product.ID).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

// Menghapus data dari database berdasarkan ID nya
func (repository *ProductRepositoryImpl) Delete(id int) error {
	if err := repository.db.Delete(&domain.Product{}, "id LIKE ?", id).Error; err != nil {
		return err
	}
	return nil
}

// Mengambil total page dari database
func (repository *ProductRepositoryImpl) GetTotalPage(model any, limit int) (totalPage int, err error) {
	var totalData int64
	if err = repository.db.Model(model).Count(&totalData).Error; err != nil {
		return 0, err
	}
	return int(totalData), err
}
