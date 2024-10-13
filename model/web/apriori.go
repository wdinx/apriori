package web

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
	ConfidanceItemset2 []ConfidanceItemset2 `json:"confidance_itemset_2"`
	ConfidanceItemset3 []ConfidanceItemset3 `json:"confidance_itemset_3"`
}
