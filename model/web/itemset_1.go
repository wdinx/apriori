package web

type ItemsetSatu struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	Count       int      `json:"count"`
	Support     float64  `json:"support"`
	Explanation string   `json:"explanation"`
}
