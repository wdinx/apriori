package service

import "apriori-backend/model/web"

type AprioriService interface {
	ProcessApriori(request *web.CreateAprioriRequest) (*web.AprioriBaseResponse, error)
	GetAll(metadata *web.Metadata) (*[]web.AprioriBaseResponse, error)
	GetByID(id string) (*web.AprioriBaseResponse, error)
	DeleteByID(id string) error
	GetRecommendationItem() (*web.RecommendationItemResponse, error)
	CreateRecommendationItem() error
}
