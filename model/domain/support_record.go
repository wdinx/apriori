package domain

type SupportRecord struct {
	ID               string             `gorm:"primaryKey;not null"`
	Itemset          string             `gorm:"type:text;not null"`
	Support          float64            `gorm:"not null"`
	AprioriID        string             `gorm:"type:varchar(255);not null"`
	OrderedStatistic []OrderedStatistic `gorm:"foreignKey:SupportRecordID"`
	AprioriResult    AprioriResult      `gorm:"foreignKey:AprioriID;references:ID"`
}
