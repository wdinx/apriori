package web

import "strings"

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
	var support float64

	if len(data.ItemsetSatu) == 0 && len(data.ItemsetDua) == 0 && len(data.ItemsetTiga) == 0 {
		data.HighestSupport = ""
		return data
	}
	if len(data.ItemsetSatu) >= 0 {
		recommendation = strings.Join(data.ItemsetSatu[0].Name, ",")
		support = data.ItemsetSatu[0].Support
	}

	if len(data.ItemsetSatu) >= 0 {
		for i := 1; i < len(data.ItemsetSatu); i++ {
			if support < data.ItemsetSatu[i].Support {
				support = data.ItemsetSatu[i].Support
				recommendation = strings.Join(data.ItemsetSatu[i].Name, ",")
			}
		}
	}

	if len(data.ItemsetDua) >= 0 {
		for i := 0; i < len(data.ItemsetDua); i++ {
			if support < data.ItemsetDua[i].Support {
				support = data.ItemsetDua[i].Support
				recommendation = strings.Join(data.ItemsetDua[i].Name, ",")
			}
		}
	}

	if len(data.ItemsetTiga) >= 0 {
		for i := 0; i < len(data.ItemsetTiga); i++ {
			if support < data.ItemsetTiga[i].Support {
				support = data.ItemsetTiga[i].Support
				recommendation = strings.Join(data.ItemsetTiga[i].Name, ",")
			}
		}
	}

	data.HighestSupport = recommendation
	return data
}
