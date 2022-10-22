package main

import (
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
	"github.com/thirumathikart/thirumathikart-auth-service/routes"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	config.InitConfig()

	config.ConnectDB()
	models.MigrateDB()

	server := echo.New()
	utils.InitLogger(server)
	server.Use(middleware.CORS())

	routes.Init(server)

	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
