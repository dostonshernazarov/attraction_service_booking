package repo

import (
	pba "attraction_service_booking/genproto/attraction_proto"
)

// UserStorageI ...
type AttractionStorageI interface {
	CreateOwner(req *pba.Owner) (*pba.Owner, error)
	GetOwner(req *pba.GetOwnerRequest) (*pba.GetOwnerResponse, error)
	DeleteOwner(req *pba.DeleteOwnerRequest) (*pba.DeleteOwnerResponse, error)
	UpdateOwnerImage(req *pba.UpdateOwnerImageRequest) (*pba.UpdateOwnerImageResponse, error)
	GetAllOwners(req *pba.GetAllOwnersRequest) (*pba.GetAllOwnersResponse, error)
	CreateAttraction(req *pba.Attraction) (*pba.Attraction, error)
	GetAttractionByName(req *pba.GetAttractionByNameRequest) (*pba.GetAttractionByNameResponse, error)
	GetAttractionsByCategory(req *pba.GetAttractionsByCategoryRequest) (*pba.GetAttractionsByCategoryResponse, error)
}
