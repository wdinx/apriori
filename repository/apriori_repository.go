package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
)

type AprioriRepository interface {
	FindAll(metadata *web.Metadata) (*[]domain.AprioriData, error)
	Create(apriori *domain.AprioriData) error
	GetByID(id string) (*domain.AprioriData, error)
	Delete(id string) error
}
