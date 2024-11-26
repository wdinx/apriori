package domain

import (
	"apriori-backend/model/web"
	"fmt"
	"gorm.io/gorm"
)

type RecommendationItem struct {
	gorm.Model
	Name string `json:"name"`
}

func (r *RecommendationItem) ToResponse(data string) *web.RecommendationItemResponse {
	return &web.RecommendationItemResponse{
		Message: fmt.Sprintf("Item yang memiliki nilai support tertinggi adalah %s", data),
	}
}
