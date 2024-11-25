package domain

type RuleAssociation struct {
	ID            string  `gorm:"primaryKey;not null"`
	Name          string  `gorm:"type:text;not null"`
	Confidence    float64 `gorm:"not null"`
	LiftRatio     float64 `gorm:"not null"`
	Explanation   string  `gorm:"type:text;not null"`
	AprioriDataID string  `gorm:"size:191;not null"`
}
