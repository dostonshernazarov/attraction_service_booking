package storage

import (
	"attraction_service_booking/storage/postgres"
	"attraction_service_booking/storage/repo"

	"github.com/jmoiron/sqlx"
)

// IStorage ...
type IStorage interface {
	Attraction() repo.AttractionStorageI
}

type Pg struct {
	db       *sqlx.DB
	userRepo repo.AttractionStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *Pg {
	return &Pg{
		db:       db,
		userRepo: postgres.NewAttractionRepo(db),
	}
}

func (s Pg) Attraction() repo.AttractionStorageI {
	return s.userRepo
}
