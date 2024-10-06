package web

type Recommendation struct {
	ID                 int `gorm:"primaryKey;autoIncrement;not null"`
	AprioriID          int
	Product            string `gorm:"type:text;not null"`
	RecommendToBuyWith string `gorm:"type:text;not null"`
}
