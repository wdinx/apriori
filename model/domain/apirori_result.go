package domain

// Struct Hasil Apriori
type AprioriResult struct {
	ID            string          `gorm:"primaryKey;not null"`
	DateStart     string          `gorm:"type:text;not null"`
	DateEnd       string          `gorm:"type:text;not null"`
	MinSupport    float64         `gorm:"not null"`
	MinConfidence float64         `gorm:"not null"`
	SupportRecord []SupportRecord `gorm:"foreignKey:AprioriID"`
}
