package web

import (
	"fmt"
	"strconv"
	"strings"
)

type CreateAprioriRequest struct {
	DateStart string  `json:"date_start" form:"date_start" validate:"required"`
	DateEnd   string  `json:"date_end" form:"date_end" validate:"required"`
	MinSup    float64 `json:"min_sup" form:"min_sup" validate:"required"`
	MinConf   float64 `json:"min_conf" form:"min_conf" validate:"required"`
}

type AprioriResponse struct {
	Items            []string           `json:"items"`
	Support          float64            `json:"support"`
	OrderedStatistic []OrderedStatistic `json:"ordered_statistic"`
}

type OrderedStatistic struct {
	Base       []string `json:"base"`
	Add        []string `json:"add"`
	Confidence float64  `json:"confidence"`
	Lift       float64  `json:"lift"`
}

//type AprioriBaseResponse struct {
//	FrequentItemset *[]FrequentItemset `json:"frequent_itemset"`
//	AssociationRule *[]AssociationRule `json:"association_rule"`
//	Recommendation  *[]Recommendation  `json:"recommendation"`
//	DecisionSupport *DecisionSupport   `json:"decision_support"`
//}

type AprioriBaseResponse struct {
	ID                 string               `json:"id"`
	DateStart          string               `json:"date_start"`
	DateEnd            string               `json:"date_end"`
	MinSupport         float64              `json:"min_support"`
	MinConfidence      float64              `json:"min_confidence"`
	ItemsetSatu        []ItemsetSatu        `json:"itemset_satu"`
	ItemsetDua         []ItemsetDua         `json:"itemset_dua"`
	ItemsetTiga        []ItemsetTiga        `json:"itemset_tiga"`
	RuleAssociation    []RuleAssociation    `json:"rule_association"`
	ConfidenceItemset2 []ConfidenceItemset2 `json:"confidance_itemset_2"`
	ConfidenceItemset3 []ConfidenceItemset3 `json:"confidance_itemset_3"`
	HighestSupport     string               `json:"highest_support"`
}

func (r *AprioriBaseResponse) GetRecommendation(data *AprioriBaseResponse) *AprioriBaseResponse {
	var recommendation string
	var support int

	if len(data.ItemsetSatu) == 0 && len(data.ItemsetDua) == 0 && len(data.ItemsetTiga) == 0 {
		data.HighestSupport = ""
		return data
	}
	if len(data.ItemsetSatu) >= 0 {
		recommendation = strings.Join(data.ItemsetSatu[0].Name, ",")
		support = ParsePercent(data.ItemsetSatu[0].Support)
	}

	if len(data.ItemsetSatu) >= 1 {
		for i := 1; i < len(data.ItemsetSatu); i++ {
			newData := ParsePercent(data.ItemsetSatu[i].Support)
			if support < newData {
				support = newData
				recommendation = strings.Join(data.ItemsetSatu[i].Name, ",")
			} else if support == newData {
				newString := strings.Join(data.ItemsetSatu[i].Name, ",")
				recommendation = fmt.Sprintf("%s, %s", recommendation, newString)
			}
		}
	}

	if len(data.ItemsetDua) >= 0 {
		for i := 0; i < len(data.ItemsetDua); i++ {
			newData := ParsePercent(data.ItemsetDua[i].Support)
			if support < newData {
				support = newData
				recommendation = strings.Join(data.ItemsetDua[i].Name, ",")
			} else if support == newData {
				newString := strings.Join(data.ItemsetDua[i].Name, ",")
				recommendation = fmt.Sprintf("%s, %s", recommendation, newString)
			}
		}
	}

	if len(data.ItemsetTiga) >= 0 {
		for i := 0; i < len(data.ItemsetTiga); i++ {
			newData := ParsePercent(data.ItemsetTiga[i].Support)
			if support < newData {
				support = newData
				recommendation = strings.Join(data.ItemsetTiga[i].Name, ",")
			} else if support == newData {
				newString := strings.Join(data.ItemsetTiga[i].Name, ",")
				recommendation = fmt.Sprintf("%s, %s", recommendation, newString)
			}
		}
	}

	data.HighestSupport = recommendation
	return data
}

func ParsePercent(s string) int {
	trimmed := strings.TrimSuffix(s, "%")

	value, _ := strconv.Atoi(trimmed)

	return value
}
