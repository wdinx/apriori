package web

type ItemsetDua struct {
	ID          string   `json:"id"`
	Name        []string `json:"name"`
	Count       int      `json:"count"`
	Support     string   `json:"support_a"`
	Explanation string   `json:"explanation"`
}
