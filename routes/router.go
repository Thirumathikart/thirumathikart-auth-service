package routes

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/thirumathikart/thirumathikart-auth-service/config"
	controllers "github.com/thirumathikart/thirumathikart-auth-service/controllers"
	"github.com/thirumathikart/thirumathikart-auth-service/generated/user"
	"github.com/thirumathikart/thirumathikart-auth-service/middlewares"
	"github.com/thirumathikart/thirumathikart-auth-service/rpcs"
	"github.com/thirumathikart/thirumathikart-auth-service/utils"
	"google.golang.org/grpc"
)

func Init() {
	// Static files
	var wg sync.WaitGroup
	wg.Add(2)

	e := echo.New()
	utils.InitLogger(e)
	e.Use(middleware.CORS())

	userE := e.Group("/api/user")
	httpPort := config.ServerPort
	userE.POST("/signup", controllers.SignupUser)
	userE.POST("/loginCustomer", controllers.LoginCustomer)
	userE.POST("/addAddress", controllers.AddAddress)
	userE.POST("/updateAddress", controllers.UpdateAddress)
	userE.POST("/address", controllers.FetchAddress)
	userE.POST("/fcmRegister", controllers.FcmRegistration)
	userE.POST("/fcmToken", controllers.FetchFCMToken)
	go func() {
		e.Logger.Fatal(e.Start(":" + httpPort))
	}()
	// GRPC server
	grpcPort := config.RPCPort
	grpcServer := grpc.NewServer(middlewares.WithServerUnaryInterceptor())
	user.RegisterUserServiceServer(grpcServer, &rpcs.UserRPCServer{})
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
		if err != nil {
			log.Panic("grpc server running error on", err)
		}
		err1 := grpcServer.Serve(lis)
		if err1 != nil {
			log.Panic("grpc server running error on", err1)
		}
	}()

	wg.Wait()
}
