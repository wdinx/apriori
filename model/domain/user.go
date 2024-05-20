package domain

type User struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement;not null"`
	Username string `json:"username" gorm:"unique;not null;type:varchar(255)"`
	Password string `json:"password" gorm:"not null;type:varchar(255)"`
}
