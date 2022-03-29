package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"bisadonasi/handler"
	"bisadonasi/user"
)

func main() {
	dsn := "root:12341234@tcp(127.0.0.1:3306)/bisadonasi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error, cannot connect to db ===>", err.Error())
	}

	fmt.Println("connected to database", db)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	api := r.Group("/api/v1")

	api.POST("/user/register", userHandler.RegisterUser)

	// r.GET("/users", handler)
	r.Run()

}
