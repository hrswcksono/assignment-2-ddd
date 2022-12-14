package database

import (
	"fmt"
	"hrswcksono/assignment2/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "12345"
	dbPort   = "5432"
	dbName   = "order_by"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	err = db.Debug().AutoMigrate(entity.Order{}, entity.Item{})
	if err != nil {
		panic("Migration failed: " + err.Error())
	}

}
func GetDB() *gorm.DB {
	return db
}
