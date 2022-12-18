package controllers

import (
	"context"

	"github.com/thirumathikart/thirumathikart-auth-service/config"
	"github.com/thirumathikart/thirumathikart-auth-service/rpcs/user"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
)

type UserRPCServer struct {
	user.UnimplementedUserServiceServer
}

func (UserRPCServer) AuthorizeRPC(ctx context.Context, request *user.AuthRequest) (*user.AuthResponse, error) {
	userDetails, err := utils.GetCurrentUserFromToken(request.UserToken, config.GetDB())
	if err != nil {
		return &user.AuthResponse{
			IsSuccess: false,
		}, nil
	}

	return &user.AuthResponse{
		IsSuccess: true,
		User: &user.User{
			Email:   userDetails.Email,
			Contact: userDetails.Contactno,
		}}, err
}
