package web

type CrossSellingRecommendation struct {
	ID                 int `gorm:"primaryKey;autoIncrement;not null"`
	AprioriID          int
	Product            string `gorm:"type:text;not null"`
	RecommendationWith string `gorm:"type:text;not null"`
}
