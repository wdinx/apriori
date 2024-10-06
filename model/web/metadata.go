package web

type Metadata struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
}

func (metadata *Metadata) Offset() int {
	return (metadata.Page - 1) * metadata.Limit
}
