package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	jwtToken, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyfQ.q_DQZ5GQyqkIC8mQjWhEXZOMOOBywYjf8xog-nWUBUM")
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println("ERROR")
		fmt.Println("ERROR")
	}

	if jwtToken.Valid {
		fmt.Println("VALID")
		fmt.Println("VALID")
		fmt.Println("VALID")
	} else {

		fmt.Println("INVALID")
		fmt.Println("INVALID")
		fmt.Println("INVALID")
	}
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser) //di contoh users
	api.POST("/login", userHandler.Login)           //di contoh sessions
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	router.Run()
}
