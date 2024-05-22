package domain

import "time"

type Transaction struct {
	ID    int       `gorm:"primaryKey;autoIncrement;not null"`
	Date  time.Time `gorm:"not null;type:date"`
	Items string    `gorm:"not null;type:text"`
}
