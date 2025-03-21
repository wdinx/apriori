package service

import (
	"apriori-backend/model/web"
	"apriori-backend/repository"
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

// Melakukan request ke repository untuk mengambil semua data produk dan mengirimnya ke kontroller
func (service *ProductServiceImpl) GetAll(metadata *web.Metadata) (*[]web.ProductResponse, error) {
	products, err := service.productRepository.GetAll(metadata)
	if err != nil {
		return nil, err
	}

	var response []web.ProductResponse
	for _, product := range *products {
		response = append(response, *converter.ToProductResponse(&product))
	}

	return &response, nil
}

// Melakukan request ke repository untuk mengambil data produk berdasarkan ID nya dan mengirimnya ke kontroller
func (service *ProductServiceImpl) GetByID(id int) (*web.ProductResponse, error) {
	product, err := service.productRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return converter.ToProductResponse(product), nil
}

// Melakukan request ke repository untuk menyimpan data produk dan mengirimnya ke kontroller
func (service *ProductServiceImpl) Create(request *web.ProductCreateRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	filename, err := service.imageService.UploadImage(request.Image)
	if err != nil {
		return err
	}
	product := converter.CreateToProductModel(request, filename)

	if err := service.productRepository.Create(product); err != nil {
		return err
	}
	return nil
}

// Melakukan request ke repository untuk update data produk dan mengirimnya ke kontroller
func (service *ProductServiceImpl) Update(request *web.ProductUpdateRequest) error {
	if err := service.validator.Struct(request); err != nil {
		return err
	}
	_, err := service.productRepository.GetByID(request.ID)
	if err != nil {
		return err
	}

	filename, err := service.imageService.UploadImage(request.Image)
	if err != nil {
		return err
	}
	product := converter.UpdateToProductModel(request, filename)

	if err := service.productRepository.Update(product); err != nil {
		return err
	}
	return nil
}

// Melakukan request ke repository untuk delete data produk dan mengirimnya ke kontroller
func (service *ProductServiceImpl) Delete(id int) error {
	if err := service.productRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
