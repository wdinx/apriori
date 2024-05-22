package web

type CreateAprioriRequest struct {
	DateStart string  `json:"date_start" form:"date_start" validate:"required"`
	DateEnd   string  `json:"date_end" form:"date_end" validate:"required"`
	MinSup    float64 `json:"min_sup" form:"min_sup" validate:"required"`
	MinConf   float64 `json:"min_conf" form:"min_conf" validate:"required"`
}

type AprioriResponse struct {
	Items            []string           `json:"items"`
	Support          float64            `json:"support"`
	OrderedStatistic []OrderedStatistic `json:"ordered_statistic"`
}

type OrderedStatistic struct {
	Base       []string `json:"base"`
	Add        []string `json:"add"`
	Confidence float64  `json:"confidence"`
	Lift       float64  `json:"lift"`
}
