package repository

import (
	"apriori-backend/model/domain"

	"gorm.io/gorm"
)

type RecommendationRepositoryImpl struct {
	db *gorm.DB
}

func NewRecommendationRepository(db *gorm.DB) RecommendationRepository {
	return &RecommendationRepositoryImpl{db: db}
}

// Menyimpan data item rekomendasi ke dalam database
func (r *RecommendationRepositoryImpl) Create(recommendationItem *domain.RecommendationItem) error {
	if err := r.db.Create(recommendationItem).Error; err != nil {
		return err
	}
	return nil
}

// Mengambil data yang terakhir tersimpan di database
func (r *RecommendationRepositoryImpl) GetLast() (*domain.RecommendationItem, error) {
	var recommendationItem domain.RecommendationItem
	if err := r.db.Last(&recommendationItem).Error; err != nil {
		return nil, err
	}
	return &recommendationItem, nil
}
