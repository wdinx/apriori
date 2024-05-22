package service

import "apriori-backend/model/web"

type AprioriService interface {
	GetApriori(request *web.CreateAprioriRequest) (*[]web.AprioriResponse, error)
}
