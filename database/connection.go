package database

import (
	"SideProject/Go-React-Demo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:Elijah_0915@/go_react"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	connection.AutoMigrate(&models.User{})
}
