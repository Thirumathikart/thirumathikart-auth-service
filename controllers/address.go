package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
)


type AddAddressRequest struct {
	Line1     string 	`json:"line1"`
	Line2     string 	`json:"line2"`
	Landmark  string 	`json:"landmark"`
	District  string 	`json:"district"`
	State     string 	`json:"state"`
	Pincode   string 	`json:"pincode"`
	Latitude  float64 	`json:"latitude"`
	Longitude float64 	`json:"longitude"`
	Token     string 	`json:"user_token"`
}

type AddAddressResponse struct {
	Message   string `json:"message"`
}

type UpdateAddressRequest struct {
	AddressId uint		`json:"id"`
	Line1     string 	`json:"line1"`
	Line2     string 	`json:"line2"`
	Landmark  string 	`json:"landmark"`
	District  string 	`json:"district"`
	State     string 	`json:"state"`
	Pincode   string 	`json:"pincode"`
	Latitude  float64 	`json:"latitude"`
	Longitude float64 	`json:"longitude"`
	Token     string 	`json:"user_token"`
}

type UpdateAddressResponse struct {
	Message   string `json:"message"`
}

type AddressResponse struct {
	UserId	  uint		`json:"userId"`
	AddressId uint		`json:"id"`
	Line1     string 	`json:"line1"`
	Line2     string 	`json:"line2"`
	Landmark  string 	`json:"landmark"`
	District  string 	`json:"district"`
	State     string 	`json:"state"`
	Pincode   string 	`json:"pincode"`
	Latitude  float64 	`json:"latitude"`
	Longitude float64 	`json:"longitude"`
}

type FetchAddressResponse struct {
	Address   []AddressResponse `json:"address"`
	Message   string `json:"message"`
}


func AddAddress(c echo.Context) error {
	var req AddAddressRequest
	//Get User
	var user models.User
	user,err:=utils.GetCurrentUser(c)
	if err!=nil{
		return utils.SendResponse(c, http.StatusUnauthorized, AddAddressResponse{Message: "UnAuthorized"})
	}

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, AddAddressResponse{Message: "Invalid Request"})
	}

	db := config.GetDB()
	address := models.Address{
		UserID:   	user.ID,
		Line1:   	req.Line1,
		Line2:    	req.Line2,
		Landmark: 	req.Landmark,
		District: 	req.District,
		State:    	req.State,
		Pincode:  	req.Pincode,
		Latitude: 	req.Latitude,
		Longitude: 	req.Longitude,
	}

	//Save Address
	db.Create(&address)

	return utils.SendResponse(c, http.StatusOK, AddAddressResponse{Message: "success"})
}


func UpdateAddress(c echo.Context) error {
	var req UpdateAddressRequest
	//Get User
	var user models.User
	user,err:=utils.GetCurrentUser(c)
	if err!=nil{
		return utils.SendResponse(c, http.StatusUnauthorized, UpdateAddressResponse{Message: "UnAuthorized"})
	}

	if err := c.Bind(&req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, UpdateAddressResponse{Message: "Invalid Request"})
	}

	db := config.GetDB()
	var address models.Address
	
	if err:= db.First(&address,"id = ?",req.AddressId).Error; err!=nil{
		return utils.SendResponse(c,http.StatusInternalServerError,"Internal Server Error")
	}
	if address.UserID != user.ID{
		return utils.SendResponse(c, http.StatusBadRequest, UpdateAddressResponse{Message: "Invalid Request"})
	}

	address.Line1 = req.Line1
	address.Line2 = req.Line2
	address.Landmark =req.Landmark
	address.District=req.District
	address.State=req.State
	address.Pincode=req.Pincode
	address.Latitude=req.Latitude
	address.Longitude=req.Longitude
	db.Save(&address)

	return utils.SendResponse(c, http.StatusOK, UpdateAddressResponse{Message: "success"})
}

func FetchAddress(c echo.Context) error {
	//Get User
	var user models.User
	user,err:=utils.GetCurrentUser(c)
	if err!=nil{
		return utils.SendResponse(c, http.StatusUnauthorized, FetchAddressResponse{Message: err.Error()})
	}

	var addresses []models.Address

	db := config.GetDB()
	if err:= db.Find(&addresses,"user_id = ?",user.ID).Error; err != nil{
		return utils.SendResponse(c,http.StatusInternalServerError,FetchAddressResponse{Message:"Internal Server Error"})
	}
	var response []AddressResponse

	for _, addr := range addresses {
		response = append(response, AddressResponse{
			UserId: 	addr.UserID,
			AddressId: 	addr.ID,
			Line1:   	addr.Line1,
			Line2:    	addr.Line2,
			Landmark: 	addr.Landmark,
			District: 	addr.District,
			State:    	addr.State,
			Pincode:  	addr.Pincode,
			Latitude: 	addr.Latitude,
			Longitude: 	addr.Longitude,
		})
	}

	return utils.SendResponse(c,http.StatusOK,FetchAddressResponse{Address: response,Message:"success"})

}


