package domain

import (
	"apriori-backend/model/web"
	Apriori "github.com/eMAGTechLabs/go-apriori"
	"github.com/google/uuid"
	"strings"
)

type AprioriResult struct {
	ID            string          `gorm:"primaryKey;not null"`
	DateStart     string          `gorm:"type:text;not null"`
	DateEnd       string          `gorm:"type:text;not null"`
	MinSupport    float64         `gorm:"not null"`
	MinConfidence float64         `gorm:"not null"`
	SupportRecord []SupportRecord `gorm:"foreignKey:AprioriID"`
}

func (a *AprioriResult) ProceedData(apriori []Apriori.RelationRecord, request *web.CreateAprioriRequest) *AprioriResult {
	a.ID = uuid.New().String()
	a.DateStart = request.DateStart
	a.DateEnd = request.DateEnd
	a.MinSupport = request.MinSup
	a.MinConfidence = request.MinConf
	var supportRecordID, orderedStatisticID uuid.UUID
	for _, record := range apriori {
		if len(record.GetSupportRecord().GetItems()) > 3 {
			continue
		}
		supportRecordID = uuid.New()
		a.SupportRecord = append(a.SupportRecord, SupportRecord{
			ID:        supportRecordID.String(),
			Itemset:   strings.Join(record.GetSupportRecord().GetItems(), ","),
			Support:   record.GetSupportRecord().GetSupport(),
			AprioriID: a.ID,
		})
		for _, statistic := range record.GetOrderedStatistic() {
			orderedStatisticID = uuid.New()
			a.SupportRecord[len(a.SupportRecord)-1].OrderedStatistic = append(a.SupportRecord[len(a.SupportRecord)-1].OrderedStatistic, OrderedStatistic{
				ID:              orderedStatisticID.String(),
				Antecedent:      strings.Join(statistic.GetBase(), ","),
				Consequent:      strings.Join(statistic.GetAdd(), ","),
				Confident:       statistic.GetConfidence(),
				Lift:            statistic.GetLift(),
				SupportRecordID: supportRecordID.String(),
			})
		}
	}
	return a
}

func (a *AprioriResult) ToResponse() *web.AprioriBaseResponse {
	var response web.AprioriBaseResponse

	response.ID = a.ID
	response.DateStart = a.DateStart
	response.DateEnd = a.DateEnd
	response.MinSupport = a.MinSupport
	response.MinConfidence = a.MinConfidence

	for _, data := range a.SupportRecord {
		itemset := strings.Split(data.Itemset, ",")
		switch {
		case len(itemset) == 1:
			response.ItemsetSatu = append(response.ItemsetSatu, web.ItemsetSatu{
				Name:        itemset,
				Count:       0,
				Support:     data.Support,
				Explanation: "Lolos",
			})
		case len(itemset) == 2:
			for _, dataDuaSet := range data.OrderedStatistic {
				antecedent := strings.Split(dataDuaSet.Antecedent, ",")
				antecedent = append(antecedent, dataDuaSet.Consequent)
				response.RuleAssociation = append(response.RuleAssociation, web.RuleAssociation{
					Name:       dataDuaSet.Antecedent + " -> " + dataDuaSet.Consequent,
					Confidence: dataDuaSet.Confident,
					LiftRatio:  dataDuaSet.Lift,
					Explanation: func() string {
						if dataDuaSet.Lift >= 1 {
							return "Korelasi Positif"
						}
						return "Korelasi Negatif"
					}()},
				)
				response.ConfidenceItemset2 = append(response.ConfidenceItemset2, web.ConfidenceItemset2{
					Name:        antecedent,
					Support:     data.Support,
					Confidence:  dataDuaSet.Confident,
					Explanation: "Lolos",
				})
			}
			response.ItemsetDua = append(response.ItemsetDua, web.ItemsetDua{
				Name:        itemset,
				Count:       0,
				Support:     data.Support,
				Explanation: "Lolos",
			})
		case len(itemset) == 3:
			for _, dataTigaSet := range data.OrderedStatistic {
				antecedent := strings.Split(dataTigaSet.Antecedent, ",")
				antecedent = append(antecedent, dataTigaSet.Consequent)
				response.RuleAssociation = append(response.RuleAssociation, web.RuleAssociation{
					Name:       dataTigaSet.Antecedent + " -> " + dataTigaSet.Consequent,
					Confidence: dataTigaSet.Confident,
					LiftRatio:  dataTigaSet.Lift,
					Explanation: func() string {
						if dataTigaSet.Lift >= 1 {
							return "Korelasi Positif"
						}
						return "Korelasi Negatif"
					}()},
				)
				response.ConfidenceItemset3 = append(response.ConfidenceItemset3, web.ConfidenceItemset3{
					Name:        antecedent,
					Support:     data.Support,
					Confidence:  dataTigaSet.Confident,
					Explanation: "Lolos",
				})
			}
			response.ItemsetTiga = append(response.ItemsetTiga, web.ItemsetTiga{
				Name:        itemset,
				Count:       0,
				Support:     data.Support,
				Explanation: "Lolos",
			})
		default:
			continue

		}
	}
	return &response
}
