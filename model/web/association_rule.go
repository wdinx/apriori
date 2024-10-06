package web

type AssociationRule struct {
	ID         int     `gorm:"primaryKey;autoIncrement;not null"`
	AprioriID  int     `gorm:"not null"`
	Antecedent string  `gorm:"type:text;not null"`
	Consequent string  `gorm:"type:text;not null"`
	Confidence float64 `gorm:"not null"`
	Lift       float64 `gorm:"not null"`
}
