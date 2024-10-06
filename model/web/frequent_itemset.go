package web

type FrequentItemset struct {
	ID        int `gorm:"primaryKey;autoIncrement;not null"`
	AprioriID int
	Itemset   string  `gorm:"type:text;not null"`
	Support   float64 `gorm:"not null"`
}
