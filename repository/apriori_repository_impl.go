package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AprioriRepositoryImpl struct {
	db *gorm.DB
}

func NewAprioriRepository(db *gorm.DB) AprioriRepository {
	return &AprioriRepositoryImpl{db: db}
}

// Mengambil semua data hasil  apriori dari database
func (r *AprioriRepositoryImpl) FindAll(metadata *web.Metadata) (*[]domain.AprioriData, error) {
	var apriori []domain.AprioriData

	if err := r.db.Order("created_at DESC").Limit(metadata.Limit).Offset(metadata.Offset()).Find(&apriori).Error; err != nil {
		return &[]domain.AprioriData{}, errors.New("Error When Fetching Apriori Result")
	}
	fmt.Println(metadata.TotalItem)
	return &apriori, nil
}

// Memasukkan atau menyimpan data hasil apriori ke dalam database
func (r *AprioriRepositoryImpl) Create(apriori *domain.AprioriData) error {
	if err := r.db.Create(apriori).Error; err != nil {
		return errors.New("Error When Saving Apriori Result")
	}
	return nil
}

// Mengambil data hasil apriori berdasarkan ID nya di database
func (r *AprioriRepositoryImpl) GetByID(id string) (*domain.AprioriData, error) {
	var apriori domain.AprioriData
	if err := r.db.Preload("ItemsetSatu").Preload("ItemsetDua").Preload("ConfidenceItemset2").Preload("ItemsetTiga").Preload("ConfidenceItemset3").Preload("RuleAssociation").First(&apriori, "id LIKE ?", id).Error; err != nil {
		return nil, errors.New("Error When Fetching Apriori Result")
	}
	return &apriori, nil
}

// Menghapus data hasil apriori dari database berdasarkan ID nya.
func (r *AprioriRepositoryImpl) Delete(id string) error {
	if err := r.db.Delete(&domain.AprioriData{}, "id LIKE ?", id).Error; err != nil {
		return errors.New("Error When Deleting Apriori Result")
	}
	return nil
}
