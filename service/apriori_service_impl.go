package service

import (
	"apriori-backend/model/domain"
	"apriori-backend/model/web"
	"apriori-backend/repository"
	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/go-playground/validator/v10"
)

type AprioriServiceImpl struct {
	transactionRepository repository.TransactionRepository
	validator             *validator.Validate
}

func NewAprioriService(transactionRepository repository.TransactionRepository, validator *validator.Validate) AprioriService {
	return &AprioriServiceImpl{
		transactionRepository: transactionRepository,
		validator:             validator,
	}
}

func (service *AprioriServiceImpl) GetApriori(request *web.CreateAprioriRequest) (*web.AprioriBaseResponse, error) {
	//if err := service.validator.Struct(request); err != nil {
	//	return nil, err
	//}
	//result, err := service.transactionRepository.FindByDateRange(request.DateStart, request.DateEnd)
	//if err != nil {
	//	return nil, err
	//}
	//var transaction [][]string
	//
	//for _, column := range *result {
	//	newColumn := strings.Split(column.Items, ",")
	//	for _, value := range newColumn {
	//		split := strings.Split(value, ",")
	//		transaction = append(transaction, split)
	//	}
	//}
	//option := Apriori.NewOptions(request.MinSup, request.MinConf, 0., 0)
	//apriori := Apriori.NewApriori(transaction[1:])
	//aprioriResult := apriori.Calculate(option)

	transactions := [][]string{
		{"beer", "nuts", "cheese"},
		{"beer", "nuts", "jam"},
		{"beer", "butter"},
		{"nuts", "cheese"},
		{"beer", "nuts", "cheese", "jam"},
		{"butter"},
		{"beer", "nuts", "jam", "butter"},
		{"jam"},
	}

	apriori := Apriori.NewApriori(transactions)
	aprioriResult := apriori.Calculate(Apriori.NewOptions(0.6, 0.6, 0.0, 0))

	var aprioriData domain.AprioriResult
	proceedApriori := aprioriData.ProceedData(aprioriResult)
	aprioriResponse := proceedApriori.ToResponse()
	return aprioriResponse, nil
}
