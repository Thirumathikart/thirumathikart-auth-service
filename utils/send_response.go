package utils

import "github.com/labstack/echo/v4"


func SendResponse(c echo.Context, code int, body interface{}) error {	
	return c.JSON(code, body)
}
