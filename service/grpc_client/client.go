package grpcClient

import (
	"attraction_service_booking/config"
	pbu "attraction_service_booking/genproto/user_proto"
	"fmt"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
}

type serviceManager struct {
	cfg         config.Config
	userService pbu.UserServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dial to user-service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.HotelServiceHost, cfg.HotelServicePort)
	}
	return &serviceManager{
		cfg:         cfg,
		userService: pbu.NewUserServiceClient(connUser),
	}, nil
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}
