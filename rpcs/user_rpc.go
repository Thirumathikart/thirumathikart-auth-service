package rpcs

import (
	"context"
	"log"

	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/generated/user"
	"github.com/thirumathikart/thirumathikart-auth-service/models"
)

type AuthRPCServer struct {
	user.UnimplementedUserServiceServer
}

func (AuthRPCServer) UserRPC(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	var userDetails models.User
	userError := config.GetDB().First(&userDetails, "id = ?", request.UserID).Error
	log.Println("USER_DETAILS", userDetails)
	if userError != nil {
		return &user.UserResponse{}, userError
	}
	var addressDetails models.Address
	addressError := config.GetDB().First(&addressDetails, "user_id = ?", request.UserID).Error
	log.Println("ADDRESS_DETAILS", addressDetails)

	if addressError != nil {
		return &user.UserResponse{}, addressError
	}

	return &user.UserResponse{
		User: &user.User{
			UserId:    uint32(userDetails.ID),
			FcmToken:  &userDetails.FcmToken,
			FirstName: &userDetails.Firstname,
			LastName:  &userDetails.Lastname,
			Email:     userDetails.Email,
			Contact:   userDetails.Contactno,
			Address: &user.Address{
				AddressId: uint32(addressDetails.ID),
				Line_1:    &addressDetails.Line1,
				Line_2:    &addressDetails.Line2,
				Name:      userDetails.Firstname + " " + userDetails.Lastname,
				Latitude:  addressDetails.Latitude,
				Longitude: addressDetails.Longitude,
				Landmark:  &addressDetails.Landmark,
				District:  &addressDetails.District,
				State:     &addressDetails.State,
				Pincode:   &addressDetails.Pincode,
			},
		}}, nil
}