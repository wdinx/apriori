package web

type ConfidanceItemset3 struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	Support     float64  `json:"support"`
	Confidance  float64  `json:"confidance"`
	Explanation string   `json:"explanation"`
}