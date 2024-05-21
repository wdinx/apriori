package config

import (
	"apriori-backend/model/domain"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB(database Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		database.DBUser,
		database.DBPass,
		database.DBHost,
		database.DBPort,
		database.DBName,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatalf("error when connecting to the database : %s", err.Error())
	}
	log.Println("connected to the database")
	Migrate()
	return db
}

func Migrate() {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&domain.User{}, &domain.Product{})
	if err != nil {
		log.Fatalf("error migratin database: %s", err.Error())
	}
}
