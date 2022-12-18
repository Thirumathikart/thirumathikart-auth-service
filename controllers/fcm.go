package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
)

type FCMRegistrationRequest struct {
	Token    string `json:"user_token"`
	FCMToken string `json:"fcm_token"`
}

type FCMRegistrationResponse struct {
	Message string `json:"message"`
}

type FCMTokenResponse struct {
	FCMToken string `json:"fcm_token"`
	Message  string `json:"message"`
}

func FcmRegistration(c echo.Context) error {
	var req FCMRegistrationRequest
	var requestedUser models.User
	requestedUser, err := utils.GetCurrentUser(c)
	if err != nil {
		return utils.SendResponse(c, http.StatusUnauthorized, FCMRegistrationResponse{Message: "UnAuthorized User"})
	}

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, FCMRegistrationResponse{Message: "Invalid Request"})
	}

	db := config.GetDB()

	//Fetch User
	var user models.User
	if err := db.First(&user, "email = ?", requestedUser.Email).Error; err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, FCMRegistrationResponse{Message: "Invalid Request"})
	}

	user.FcmToken = req.FCMToken
	db.Save(&user)
	return utils.SendResponse(c, http.StatusOK, FCMRegistrationResponse{Message: "success"})
}

func FetchFCMToken(c echo.Context) error {
	var requestedUser models.User
	requestedUser, err := utils.GetCurrentUser(c)
	if err != nil {
		return utils.SendResponse(c, http.StatusUnauthorized, FCMTokenResponse{Message: "UnAuthorized User"})
	}

	db := config.GetDB()

	//Fetch User
	var user models.User
	if err := db.First(&user, "email = ?", requestedUser.Email).Error; err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, FCMTokenResponse{Message: "Invalid Request"})
	}

	if user.FcmToken == "" {
		return utils.SendResponse(c, http.StatusBadRequest, FCMTokenResponse{Message: "Invalid Request"})
	}

	return utils.SendResponse(c, http.StatusOK, FCMTokenResponse{FCMToken: user.FcmToken, Message: "success"})
}
