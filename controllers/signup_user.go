package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
)

type SignupRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsSeller  bool   `json:"is_seller"`
}

type SignUpResponse struct {
	Message   string `json:"message"`
}

func SignupUser(c echo.Context) error {
	var req SignupRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, SignUpResponse{Message :"Invalid request"})
	}

	// Generate Password Hash
	hasher := sha1.New()
	hasher.Write([]byte(req.Password))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))

	newUser := models.User{
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Email:        req.Email,
		IsSeller:     req.IsSeller,
		Passwordhash: passwordHash,
	}

	db := config.GetDB()

	if err := db.Create(&newUser).Error; err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, SignUpResponse{Message :"User already exists"})
	}

	return utils.SendResponse(c, http.StatusOK, SignUpResponse{Message :"User created successfully"})
}
