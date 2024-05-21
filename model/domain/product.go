package domain

type Product struct {
	ID    int    `gorm:"primaryKey;autoIncrement;not null"`
	Name  string `gorm:"not null;type:varchar(255)"`
	Price int    `gorm:"not null"`
	Image string `gorm:"not null;type:varchar(255)"`
}
