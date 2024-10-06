package web

type HighSupportProduct struct {
	ID        int `gorm:"primaryKey;autoIncrement;not null"`
	AprioriID int
	Item      string  `gorm:"type:text;not null"`
	Support   float64 `gorm:"not null"`
}
