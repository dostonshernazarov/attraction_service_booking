package repo

import (
	pba "attraction_service_booking/genproto/attraction_proto"
)

// UserStorageI ...
type AttractionStorageI interface {
	// owner
	CreateOwner(req *pba.Owner) (*pba.Owner, error)
	CheckUniqueness(req *pba.UniquenessRequest) (*pba.UniquenessResponse, error)
	GetOwner(req *pba.GetOwnerRequest) (*pba.GetOwnerResponse, error)
	DeleteOwner(req *pba.DeleteOwnerRequest) (*pba.DeleteOwnerResponse, error)
	UpdateOwnerImage(req *pba.UpdateOwnerImageRequest) (*pba.UpdateOwnerImageResponse, error)
	GetAllOwners(req *pba.GetAllOwnersRequest) (*pba.GetAllOwnersResponse, error)

	// attraction
	CreateAttraction(req *pba.Attraction) (*pba.Attraction, error)
	GetAttractionsByName(req *pba.GetAttractionByNameRequest) (*pba.GetAttractionByNameResponse, error)
	GetAttractionsByCategory(req *pba.GetAttractionsByCategoryRequest) (*pba.GetAttractionsByCategoryResponse, error)
	GetAttractionsByLocation(req *pba.GetAttractionsByLocationRequest) (*pba.GetAttractionsByLocationResponse, error)
	GetAttractionsByRating(req *pba.GetAttractionsByRatingRequest) (*pba.GetAttractionsByRatingResponse, error)
	UpdateAttractionImage(req *pba.UpdateAttractionImageRequest) (*pba.UpdateAttractionImageResponse, error)
	DeleteAttractionImage(req *pba.DeleteAttractionImageRequest) (*pba.DeleteAttractionImageResponse, error)

	// review
	AddReview(req *pba.Review) (*pba.Review, error)
	GetReview(req *pba.GetReviewRequest) (*pba.GetReviewResponse, error)
	ListReviews(req *pba.ListReviewsRequest) (*pba.ListReviewsResponse, error)
	UpdateReviewComment(req *pba.UpdateReviewCommentRequest) (*pba.UpdateReviewCommentResponse, error)
	DeleteReview(req *pba.DeleteReviewRequest) (*pba.DeleteReviewResponse, error)

	// favourites
	AddToFavourites(req *pba.AddToFavouritesRequest) (*pba.AddToFavouritesResponse, error)
	DropFromFavourites(req *pba.DropFromFavouritesRequest) (*pba.DropFromFavouritesResponse, error)
	ListOfFavourites(req *pba.ListOfFavouritesRequest) (*pba.ListOfFavouritesResponse, error)
}
