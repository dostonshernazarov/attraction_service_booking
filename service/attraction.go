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

//	****OWNER****
//
// CreateOwner implementation
func (s *AttractionService) CreateOwner(ctx context.Context, req *pba.Owner) (*pba.Owner, error) {
	owner, err := s.storage.Attraction().CreateOwner(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return owner, nil
}

// CheckUniqueness implementation
func (s *AttractionService) CheckUniqueness(ctx context.Context, req *pba.UniquenessRequest) (*pba.UniquenessResponse, error) {
	success, err := s.storage.Attraction().CheckUniqueness(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
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

//	****ATTRACTION****
//
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
func (s *AttractionService) GetAttractionsByName(ctx context.Context, req *pba.GetAttractionByNameRequest) (*pba.GetAttractionByNameResponse, error) {
	attraction, err := s.storage.Attraction().GetAttractionsByName(req)
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

// GetAttractionsByLocation implementation
func (s *AttractionService) GetAttractionsByLocation(ctx context.Context, req *pba.GetAttractionsByLocationRequest) (*pba.GetAttractionsByLocationResponse, error) {
	attractions, err := s.storage.Attraction().GetAttractionsByLocation(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return attractions, nil
}

// GetAttractionsByLocation implementation
func (s *AttractionService) UpdateAttractionImage(ctx context.Context, req *pba.UpdateAttractionImageRequest) (*pba.UpdateAttractionImageResponse, error) {
	success, err := s.storage.Attraction().UpdateAttractionImage(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

// DeleteAttractionImage implementation
func (s *AttractionService) DeleteAttractionImage(ctx context.Context, req *pba.DeleteAttractionImageRequest) (*pba.DeleteAttractionImageResponse, error) {
	success, err := s.storage.Attraction().DeleteAttractionImage(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

// GetAttractionsByRating implementation
func (s *AttractionService) GetAttractionsByRating(ctx context.Context, req *pba.GetAttractionsByRatingRequest) (*pba.GetAttractionsByRatingResponse, error) {
	success, err := s.storage.Attraction().GetAttractionsByRating(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

//	****REVIEW****

// AddReview implementation
func (s *AttractionService) AddReview(ctx context.Context, req *pba.Review) (*pba.Review, error) {
	review, err := s.storage.Attraction().AddReview(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return review, nil
}

// GetReview implementation
func (s *AttractionService) GetReview(ctx context.Context, req *pba.GetReviewRequest) (*pba.GetReviewResponse, error) {
	review, err := s.storage.Attraction().GetReview(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return review, nil
}

// ListReviews implementation
func (s *AttractionService) ListReviews(ctx context.Context, req *pba.ListReviewsRequest) (*pba.ListReviewsResponse, error) {
	reviews, err := s.storage.Attraction().ListReviews(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return reviews, nil
}

// UpdateReviewComment implementation
func (s *AttractionService) UpdateReviewComment(ctx context.Context, req *pba.UpdateReviewCommentRequest) (*pba.UpdateReviewCommentResponse, error) {
	review, err := s.storage.Attraction().UpdateReviewComment(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return review, nil
}

// DeleteReview implementation
func (s *AttractionService) DeleteReview(ctx context.Context, req *pba.DeleteReviewRequest) (*pba.DeleteReviewResponse, error) {
	success, err := s.storage.Attraction().DeleteReview(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return success, nil
}

//	****FAVOURITES****

// AddToFavourites implementation
func (s *AttractionService) AddToFavourites(ctx context.Context, req *pba.AddToFavouritesRequest) (*pba.AddToFavouritesResponse, error) {
	favourite, err := s.storage.Attraction().AddToFavourites(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return favourite, nil
}

// DropFromFavourites implementation
func (s *AttractionService) DropFromFavourites(ctx context.Context, req *pba.DropFromFavouritesRequest) (*pba.DropFromFavouritesResponse, error) {
	success, err := s.storage.Attraction().DropFromFavourites(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return success, nil
}

// ListOfFavourites implementation
func (s *AttractionService) ListOfFavourites(ctx context.Context, req *pba.ListOfFavouritesRequest) (*pba.ListOfFavouritesResponse, error) {
	favourites, err := s.storage.Attraction().ListOfFavourites(req)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return favourites, nil
}


