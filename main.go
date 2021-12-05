package main

import (
	"fmt"
	"go-startup/auth"
	"go-startup/campaign"

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
	campaignRepository := campaign.NewRepository(db)

	userService := user.NewService(userRepository)

	campaignService := campaign.NewService(campaignRepository)

	authService := auth.NewService()

	campaigns, err := campaignService.FindCampaigns(0)
	fmt.Println(len(campaigns))

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
		// mendapatkan headernya dengan key Authorization
		authHeader := c.GetHeader("Authorization")

		// ini untuk cek bearer ada atau ga di header
		if !strings.Contains(authHeader, "Bearer") {
			respose := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
			return
		}

		// ini untuk mengsplit header tadi untuk ambil tokennya saja
		// struktur headernya
		// Bearer tokentokentoken
		// di split dengan spasi " "

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
			// mengisi token string dengan array token index ke 1 yang berisi token
		}

		// disini memvalidasi token yang dapat dari split tadi
		token, err := authService.ValidateToken(tokenString)

		if err != nil {
			respose := helper.APIResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, respose)
			return

		}
		// mengambil claim pada auth
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
