package main

import (
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/routes"
)

func main() {

	config.InitConfig()
	routes.Init()

}
