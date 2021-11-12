package main

import (
	"go-startup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/go-startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userInput := user.RegisterUserInput{}
	userInput.Name = "test simpan service"
	userInput.Email = "email@email.com"
	userInput.Occupation = "idol"
	userInput.Password = "Password"

	userService.RegisterUser(userInput)
}
