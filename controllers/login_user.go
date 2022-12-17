package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
)

type CustomerLoginRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CustomerLoginResponse struct {
	Token 	  string `json:"token"`
	Message   string `json:"message"`
}

func LoginCustomer(c echo.Context) error {
	var req CustomerLoginRequest

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, CustomerLoginResponse{Message: "Invalid Request"})
	}

	db := config.GetDB()

	//Fetch User
	var user models.User
	if err := db.First(&user,"email = ?",req.Email).Error; err != nil {
		return utils.SendResponse(c, http.StatusUnauthorized, CustomerLoginResponse{Message: "Invalid Credentials"})
	}

	// Generate Password Hash
	hasher := sha1.New()
	hasher.Write([]byte(req.Password))
	passwordHash := hex.EncodeToString(hasher.Sum(nil))

	//Password Check
	if user.Passwordhash!=passwordHash {
		return utils.SendResponse(c,http.StatusUnauthorized,CustomerLoginResponse{Message: "Invalid Credentials"})
	}

	//Generate jwt
	if token, err :=  utils.CreateToken(jwt.MapClaims{
		"email":  user.Email,
		"contact" : user.Contactno,
	});err!=nil{
		return utils.SendResponse(c,http.StatusInternalServerError,"Internal Server Error")
	} else{
		return utils.SendResponse(c, http.StatusOK, CustomerLoginResponse{Token: token,Message: "User Authenticated Successfully"})
	}
}
