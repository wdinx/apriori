package web

type ConfidenceItemset3 struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	Support     string   `json:"support"`
	Confidence  string   `json:"confidance"`
	Explanation string   `json:"explanation"`
}
