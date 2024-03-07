package grpcClient

import (
	"attraction_service_booking/config"
	pbh "attraction_service_booking/genproto/hotel_proto"
	pbu "attraction_service_booking/genproto/user_proto"
	"fmt"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	HotelService() pbh.HotelServiceClient
	UserService() pbu.UserServiceClient
}

type serviceManager struct {
	cfg          config.Config
	hotelService pbh.HotelServiceClient
	userService  pbu.UserServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	// dial to hotel-service
	connHotel, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.HotelServiceHost, cfg.HotelServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.HotelServiceHost, cfg.HotelServicePort)
	}

	// dial to user-service
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dail host: %s port : %d", cfg.HotelServiceHost, cfg.HotelServicePort)
	}
	return &serviceManager{
		cfg:          cfg,
		hotelService: pbh.NewHotelServiceClient(connHotel),
		userService:  pbu.NewUserServiceClient(connUser),
	}, nil
}

func (s *serviceManager) HotelService() pbh.HotelServiceClient {
	return s.hotelService
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}
