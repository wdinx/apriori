package service

import (
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"apriori-backend/util"
	"apriori-backend/util/converter"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
	imageService      ImageService
	validator         *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, imageService ImageService, validator *validator.Validate) ProductService {
	return &ProductServiceImpl{
		productRepository: productRepository,
		imageService:      imageService,
		validator:         validator,
	}
}

func (service *ProductServiceImpl) GetAll() (*[]web.ProductResponse, error) {
	products, err := service.productRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var response []web.ProductResponse
	for _, product := range *products {
		response = append(response, *converter.ToProductResponse(&product))
	}

	return &response, nil
}

func (service *ProductServiceImpl) GetByID(id int) (*web.ProductResponse, error) {
	product, err := service.productRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return converter.ToProductResponse(product), nil
}

func (service *ProductServiceImpl) Create(request *web.ProductCreateRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	filename := util.GenerateImageName(request.Name, request.Image.Filename)
	if err := service.imageService.UploadImage(request.Image, filename); err != nil {
		return err
	}
	product := converter.CreateToProductModel(request, filename)

	if err := service.productRepository.Create(product); err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) Update(request *web.ProductUpdateRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	_, err := service.productRepository.GetByID(request.ID)
	if err != nil {
		return err
	}

	filename := util.GenerateImageName(request.Name, request.Image.Filename)
	if err := service.imageService.UploadImage(request.Image, filename); err != nil {
		return err
	}
	product := converter.UpdateToProductModel(request, filename)

	if err := service.productRepository.Update(product); err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) Delete(id int) error {
	if err := service.productRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
