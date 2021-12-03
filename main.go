package main

import (
	"go-startup/auth"
	"go-startup/handler"
	"go-startup/helper"
	"go-startup/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

	//token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxN30.9AhPIz7AyPWBJyk0sb6uFuN2QrPctGRr45n1l7jxaL4")
	//if err != nil {
	//	fmt.Println("error")
	//}

	//if token.Valid {
	//	fmt.Println("valid")
	//} else {
	//	fmt.Println("invalid")
	//}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			respose := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			respose := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
			return

		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			respose := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
			return
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			respose := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
			return
		}

		c.Set("currentUser", user)

	}
}

//ambil nilai header authorization: bearer token
// dari header auth,kita ambil nilai token nya saja
// lalu token yang di dapat kita validasi
// token valid, dapat user_id
// ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user
