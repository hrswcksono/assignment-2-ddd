package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Items struct {
	ItemId      uint `gorm:"primaryKey"`
	ItemCode    int
	Description string
	Quantity    int
	OrderFk     uint
}

type Orders struct {
	OrderId      uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Items `gorm:"foreignKey:OrderFk;references:OrderId"`
}

var (
	username = "root"
	password = ""
	host     = "localhost"
	port     = 3306
	dbName   = "orders_by"
	db       *gorm.DB
)

func InitDb() {
	// dsn := "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Debug().AutoMigrate(&Orders{}, &Items{})
}

func GetDB() *gorm.DB {
	return db
}
