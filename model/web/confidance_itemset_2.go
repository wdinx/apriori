package web

type ConfidenceItemset2 struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	Support     string   `json:"support"`
	Confidence  string   `json:"confidance"`
	Explanation string   `json:"explanation"`
}
