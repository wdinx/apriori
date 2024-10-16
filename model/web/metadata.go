package web

type Metadata struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalItem int `json:"total_item"`
}

func (metadata *Metadata) Offset() int {
	return (metadata.Page - 1) * metadata.Limit
}
