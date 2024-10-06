package web

type DecisionSupport struct {
	HighSupportProduct         []HighSupportProduct         `json:"high_support_product"`
	CrossSellingRecommendation []CrossSellingRecommendation `json:"cross_selling_recommendation"`
}
