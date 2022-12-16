package routes

import (
	controllers "github.com/thirumathikart/thirumathikart-auth-service/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.POST("/signup", controllers.SignupUser)
	user.POST("/loginCustomer", controllers.LoginCustomer)
	user.POST("/addAddress", controllers.AddAddress)
	user.POST("/updateAddress", controllers.UpdateAddress)
	user.POST("/address", controllers.FetchAddress)
	user.POST("/fcmRegister", controllers.FcmRegistration)
	user.POST("/fcmToken", controllers.FetchFCMToken)
}
