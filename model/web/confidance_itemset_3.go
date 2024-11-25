package web

type ConfidenceItemset3 struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	Support     float64  `json:"support"`
	Confidence  float64  `json:"confidance"`
	Explanation string   `json:"explanation"`
}
