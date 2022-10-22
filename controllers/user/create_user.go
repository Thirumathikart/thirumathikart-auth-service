package handlers

import (
	"net/http"

	"github.com/thirumathikart/thirumathikart-auth-service/utils"

	"github.com/thirumathikart/thirumathikart-auth-service/models"

	"github.com/thirumathikart/thirumathikart-auth-service/config"

	"github.com/labstack/echo/v4"
)

type SignupRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	RollNo   string `json:"roll_no"`
}

func SignupUser(c echo.Context) error {
	var req SignupRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid request")
	}

	newUser := models.User{
		Username: req.Username,
		Name:     req.Name,
		RollNo:   req.RollNo,
	}

	db := config.GetDB()

	if err := db.Create(&newUser).Error; err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "User already exists")
	}

	return utils.SendResponse(c, http.StatusOK, "User created successfully")
}
