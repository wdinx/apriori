package service

import (
	"apriori-backend/model/web"
	"apriori-backend/repository"
	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/go-playground/validator/v10"
	"strings"
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

func (service *AprioriServiceImpl) GetApriori(request *web.CreateAprioriRequest) (*[]web.AprioriResponse, error) {
	if err := service.validator.Struct(request); err != nil {
		return nil, err
	}
	result, err := service.transactionRepository.FindByDateRange(request.DateStart, request.DateEnd)
	if err != nil {
		return nil, err
	}
	var transaction [][]string

	for _, column := range *result {
		newColumn := strings.Split(column.Items, ",")
		for _, value := range newColumn {
			split := strings.Split(value, ",")
			transaction = append(transaction, split)
		}
	}
	option := Apriori.NewOptions(request.MinSup, request.MinConf, 0., 0)
	apriori := Apriori.NewApriori(transaction[1:])
	aprioriResult := apriori.Calculate(option)

	var response []web.AprioriResponse
	for _, record := range aprioriResult {
		items := record.GetSupportRecord().GetItems()
		support := record.GetSupportRecord().GetSupport()
		orderedStatistic := record.GetOrderedStatistic()
		var orderedStatisticResponse []web.OrderedStatistic

		for _, statistic := range orderedStatistic {
			orderedStatisticResponse = append(orderedStatisticResponse, web.OrderedStatistic{
				Base:       statistic.GetBase(),
				Add:        statistic.GetAdd(),
				Confidence: statistic.GetConfidence(),
				Lift:       statistic.GetLift(),
			})
		}

		response = append(response, web.AprioriResponse{
			Items:            items,
			Support:          support,
			OrderedStatistic: orderedStatisticResponse,
		})
	}
	return &response, err
}
