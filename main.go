package main

import (
	"dans-multi-pro-test/controller"
	"dans-multi-pro-test/database"
	"dans-multi-pro-test/middleware"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	database.Connect()
	defer database.DB.Close()

	router := gin.Default()

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)

	api := router.Group("/api").Use(middleware.Auth())
	api.GET("/ping", controller.Ping)
	api.GET("/recruitment/positions.json", controller.GetPositions)
	api.GET("/recruitment/position/:id", controller.GetJobDetail)

	router.Run(os.Getenv("PORT"))
}
