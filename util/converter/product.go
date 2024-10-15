package converter

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
)

func CreateToProductModel(product *web.ProductCreateRequest, filename string) *domain.Product {
	return &domain.Product{
		Name:  product.Name,
		Price: product.Price,
		Image: filename,
	}
}

func UpdateToProductModel(product *web.ProductUpdateRequest, filename string) *domain.Product {
	return &domain.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Image: filename,
	}
}

func ToProductResponse(product *domain.Product) *web.ProductResponse {
	return &web.ProductResponse{
		Name:  product.Name,
		Price: product.Price,
		Image: product.Image,
	}
}
