package routes

import (
	controllers "github.com/thirumathikart/thirumathikart-auth-service/controllers/user"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.POST("/signup", controllers.SignupUser)
}