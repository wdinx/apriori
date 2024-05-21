package domain

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement;not null"`
	Username string `gorm:"unique;not null;type:varchar(255)"`
	Password string `gorm:"not null;type:varchar(255)"`
}
