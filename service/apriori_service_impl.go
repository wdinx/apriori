package service

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	"fmt"
	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/go-playground/validator/v10"
	"strings"
)

type AprioriServiceImpl struct {
	transactionRepository repository.TransactionRepository
	aprioriRepository     repository.AprioriRepository
	validator             *validator.Validate
}

func NewAprioriService(transactionRepository repository.TransactionRepository, aprioriRepository repository.AprioriRepository, validator *validator.Validate) AprioriService {
	return &AprioriServiceImpl{
		transactionRepository: transactionRepository,
		aprioriRepository:     aprioriRepository,
		validator:             validator,
	}
}

func (service *AprioriServiceImpl) ProcessApriori(request *web.CreateAprioriRequest) (*web.AprioriBaseResponse, error) {
	if err := service.validator.Struct(request); err != nil {
		return nil, err
	}
	result, err := service.transactionRepository.FindByDateRange(request.DateStart, request.DateEnd)
	if err != nil {
		return nil, err
	}

	if len(*result) == 0 {
		return nil, fmt.Errorf("Data Not Found")
	}

	var transaction [][]string

	for _, column := range *result {
		newColumn := strings.Split(column.Items, ",")
		for _, value := range newColumn {
			split := strings.Split(value, ",")
			for _, s := range split {
				newColumn = append(newColumn, s)
			}
		}
		transaction = append(transaction, newColumn)
	}

	fmt.Println(transaction)
	option := Apriori.NewOptions(request.MinSup, request.MinConf, 0., 0.)
	apriori := Apriori.NewApriori(transaction[1:])
	aprioriResult := apriori.Calculate(option)

	fmt.Println(aprioriResult)

	var aprioriData domain.AprioriData
	proceedApriori := aprioriData.ProceedData(aprioriResult, request, transaction)
	if err = service.aprioriRepository.Create(proceedApriori); err != nil {
		return nil, err
	}
	aprioriResponse := proceedApriori.ToResponse()
	return aprioriResponse, nil
}

func (service *AprioriServiceImpl) GetAll(metadata *web.Metadata) (*[]web.AprioriBaseResponse, error) {
	apriori, err := service.aprioriRepository.FindAll(metadata)

	if err != nil {
		return nil, err
	}
	var response []web.AprioriBaseResponse
	for _, data := range *apriori {
		response = append(response, *data.ToResponse())
	}
	return &response, nil
}

func (service *AprioriServiceImpl) GetByID(id string) (*web.AprioriBaseResponse, error) {
	apriori, err := service.aprioriRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return apriori.ToResponse(), nil
}

func (service *AprioriServiceImpl) DeleteByID(id string) error {
	if err := service.aprioriRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
