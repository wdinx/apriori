package repository

import "apriori-backend/model/domain"

type RecommendationRepository interface {
	Create(recommendationItem *domain.RecommendationItem) error
	GetLast() (*domain.RecommendationItem, error)
}
