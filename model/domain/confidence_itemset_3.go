package domain

type ConfidenceItemset3 struct {
	ID            string  `gorm:"primaryKey;not null"`
	Name          string  `gorm:"type:text;not null"`
	Support       float64 `gorm:"not null"`
	Confidence    float64 `gorm:"not null"`
	Explanation   string  `gorm:"type:text;not null"`
	AprioriDataID string  `gorm:"size:191;not null"`
}
