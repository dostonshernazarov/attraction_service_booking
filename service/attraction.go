package service

import (
	pba "attraction_service_booking/genproto/attraction_proto"
	l "attraction_service_booking/pkg/logger"
	"attraction_service_booking/storage"
	"context"

	grpcClient "attraction_service_booking/service/grpc_client"

	"github.com/jmoiron/sqlx"
)

// UserService ...
type AttractionService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpcClient.IServiceManager
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger, client grpcClient.IServiceManager) *AttractionService {
	return &AttractionService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

// CreateOwner implementation
func (s *AttractionService) CreateOwner(ctx context.Context, req *pba.Owner) (*pba.Owner, error) {
	owner, err := s.storage.Attraction().CreateOwner(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return owner, nil
}

// GetOwner implementation
func (s *AttractionService) GetOwner(ctx context.Context, req *pba.GetOwnerRequest) (*pba.GetOwnerResponse, error) {
	owner, err := s.storage.Attraction().GetOwner(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return owner, nil
}

// DeleteOwner implementation
func (s *AttractionService) DeleteOwner(ctx context.Context, req *pba.DeleteOwnerRequest) (*pba.DeleteOwnerResponse, error) {
	success, err := s.storage.Attraction().DeleteOwner(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

// UpdateOwnerImage implementation
func (s *AttractionService) UpdateOwnerImage(ctx context.Context, req *pba.UpdateOwnerImageRequest) (*pba.UpdateOwnerImageResponse, error) {
	success, err := s.storage.Attraction().UpdateOwnerImage(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

// GetAllUsers implementation
func (s *AttractionService) GetAllOwners(ctx context.Context, req *pba.GetAllOwnersRequest) (*pba.GetAllOwnersResponse, error) {
	owners, err := s.storage.Attraction().GetAllOwners(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return owners, nil
}

// CreateAttraction implementation
func (s *AttractionService) CreateAttraction(ctx context.Context, req *pba.Attraction) (*pba.Attraction, error) {
	attraction, err := s.storage.Attraction().CreateAttraction(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return attraction, nil
}

// GetAttractionByName implementation
func (s *AttractionService) GetAttractionByName(ctx context.Context, req *pba.GetAttractionByNameRequest) (*pba.GetAttractionByNameResponse, error) {
	attraction, err := s.storage.Attraction().GetAttractionByName(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return attraction, nil
}

// GetAttractionsByCategory implementation
func (s *AttractionService) GetAttractionsByCategory(ctx context.Context, req *pba.GetAttractionsByCategoryRequest) (*pba.GetAttractionsByCategoryResponse, error) {
	attractions, err := s.storage.Attraction().GetAttractionsByCategory(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return attractions, nil
}