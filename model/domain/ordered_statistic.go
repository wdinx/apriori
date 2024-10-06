package domain

type OrderedStatistic struct {
	ID              string  `gorm:"primaryKey;not null"`
	Antecedent      string  `gorm:"type:text;not null"`
	Consequent      string  `gorm:"type:text;not null"`
	Confident       float64 `gorm:"not null"`
	Lift            float64 `gorm:"not null"`
	SupportRecordID string  `gorm:"not null"`
}
