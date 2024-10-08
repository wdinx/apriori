package repository

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
)

type AprioriRepository interface {
	FindAll(metadata *web.Metadata) (*[]domain.AprioriResult, error)
	Create(apriori *domain.AprioriResult) error
	GetByID(id int) (*domain.AprioriResult, error)
}
