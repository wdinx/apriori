package web

type ConfidanceItemset2 struct {
	Name        []string `json:"name"`
	Support     float64  `json:"support"`
	Confidance  float64  `json:"confidance"`
	Explanation string   `json:"explanation"`
}
