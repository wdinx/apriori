package web

type RuleAssociation struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Confidence  string  `json:"confidence"`
	LiftRatio   float64 `json:"lift_ratio"`
	Explanation string  `json:"explanation"`
}
