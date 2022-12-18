package rpcs

import (
	"context"
	"log"

	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/generated/user"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
)

type UserRPCServer struct {
	user.UnimplementedUserServiceServer
}

func (UserRPCServer) AuthRPC(ctx context.Context, request *user.AuthRequest) (*user.AuthResponse, error) {
	userDetails, err := utils.GetCurrentUserFromToken(request.UserToken, config.GetDB())
	log.Println("USER_DETAILS", userDetails)

	if err != nil {
		return &user.AuthResponse{
			IsSuccess: false,
		}, err
	}

	return &user.AuthResponse{
		IsSuccess: true,
		User: &user.User{
			UserId:    uint32(userDetails.ID),
			FcmToken:  &userDetails.FcmToken,
			FirstName: &userDetails.Firstname,
			LastName:  &userDetails.Lastname,
			Email:     userDetails.Email,
			Contact:   userDetails.Contactno,
		}}, nil
}
