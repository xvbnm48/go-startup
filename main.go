package main

import (
	"fmt"
	"go-startup/auth"
	"go-startup/handler"
	"go-startup/user"
	"log"

	"github.com/gin-gonic/gin"
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
	userService.SaveAvatar(1, "images/avatar.jpg")
	authService := auth.NewService()
	//fmt.Println(authService.GenerateToken(10))

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.9AhPIz7AyPWBJyk0sb6uFuN2QrPctGRr45n1l7jxaL4")
	if err != nil {
		fmt.Println("error")
	}

	if token.Valid {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
}
