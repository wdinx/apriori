package domain

type Product struct {
	ID    int    `gorm:"primaryKey;autoIncrement;not null"`
	Name  string `gorm:"not null;type:varchar(255)"`
	Price int    `gorm:"not null"`
	Stock int    `gorm:"not null;default:1000"`
	Image string `gorm:"not null;type:varchar(255)"`
}
