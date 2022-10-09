package main

import "hrswcksono/assignment2/database"

// type User struct {
// 	gorm.Model
// 	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
// }

// type CreditCard struct {
// 	gorm.Model
// 	Number    string
// 	UserRefer uint
// }

func main() {
	database.InitDb()
	// db.AutoMigrate(&User{}, &CreditCard{})
}
