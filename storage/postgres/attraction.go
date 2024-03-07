package postgres

import (
	pba "attraction_service_booking/genproto/attraction_proto"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type AttractionRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewAttractionRepo(db *sqlx.DB) *AttractionRepo {
	return &AttractionRepo{
		db: db,
	}
}

// Create a new owner
func (r *AttractionRepo) CreateOwner(owner *pba.Owner) (*pba.Owner, error) {
	var (
		newOwner                                                                           pba.Owner
		full_name, email, password, birthday, phone_number, image_url, refresh_token, role sql.NullString
	)

	query := `INSERT INTO owner_table
	(owner_id,
	full_name,
	email,
	password,
	birthday,
	phone_number,
	image_url,
	refresh_token,
	role)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING owner_id,
	full_name,
	email,
	password,
	birthday,
	phone_number,
	image_url,
	refresh_token,
	role,
	created_at,
	updated_at,
	deleted_at`

	err := r.db.QueryRow(
		query,
		owner.OwnerId,
		sql.NullString{String: owner.FullName, Valid: owner.FullName != ""},
		sql.NullString{String: owner.Email, Valid: owner.Email != ""},
		sql.NullString{String: owner.Password, Valid: owner.Password != ""},
		sql.NullString{String: owner.Birthday, Valid: owner.Birthday != ""},
		sql.NullString{String: owner.PhoneNumber, Valid: owner.PhoneNumber != ""},
		sql.NullString{String: owner.ImageUrl, Valid: owner.ImageUrl != ""},
		sql.NullString{String: owner.RefreshToken, Valid: owner.RefreshToken != ""},
		sql.NullString{String: owner.Role, Valid: owner.Role != ""},
	).Scan(
		&newOwner.OwnerId,
		&full_name,
		&email,
		&password,
		&birthday,
		&phone_number,
		&image_url,
		&refresh_token,
		&role,
		&newOwner.CreatedAt,
		&newOwner.UpdatedAt,
		&newOwner.DeletedAt)

	if err != nil {
		return nil, err
	}

	// checking for sql.NullString
	newOwner.FullName = stringValue(full_name)
	newOwner.Email = stringValue(email)
	newOwner.Password = stringValue(password)
	newOwner.Birthday = stringValue(birthday)
	newOwner.PhoneNumber = stringValue(phone_number)
	newOwner.ImageUrl = stringValue(image_url)
	newOwner.RefreshToken = stringValue(refresh_token)
	newOwner.Role = stringValue(role)

	return &newOwner, nil
}

// Get owner by phone number
func (r *AttractionRepo) GetOwner(owner *pba.GetOwnerRequest) (*pba.GetOwnerResponse, error) {
	var getOwner pba.GetOwnerResponse

	query := `SELECT
	owner_id,
	full_name,
	email,
	password,
	birthday,
	phone_number,
	image_url,
	refresh_token,
	created_at,
	updated_at,
	deleted_at,
	role
	FROM owner_table
	WHERE email = $1`

	err := r.db.QueryRow(
		query,
		owner.PhoneNumber,
	).Scan(
		&getOwner.Owner.OwnerId,
		&getOwner.Owner.FullName,
		&getOwner.Owner.Email,
		&getOwner.Owner.Password,
		&getOwner.Owner.Birthday,
		&getOwner.Owner.ImageUrl,
		&getOwner.Owner.RefreshToken,
		&getOwner.Owner.CreatedAt,
		&getOwner.Owner.UpdatedAt,
		&getOwner.Owner.DeletedAt,
		&getOwner.Owner.Role,
	)

	if err != nil {
		return nil, err
	}

	return &getOwner, nil
}

// Delete owner by phone number
func (r *AttractionRepo) DeleteOwner(req *pba.DeleteOwnerRequest) (*pba.DeleteOwnerResponse, error) {
	query := `DELETE FROM owner_table WHERE phone_number = $1`

	_, err := r.db.Exec(query, req.PhoneNumber)
	if err != nil {
		return &pba.DeleteOwnerResponse{
			Success: false,
		}, err
	}

	return &pba.DeleteOwnerResponse{
		Success: true,
	}, nil
}

// Update owner image by phone number
func (r *AttractionRepo) UpdateOwnerImage(req *pba.UpdateOwnerImageRequest) (*pba.UpdateOwnerImageResponse, error) {
	query := `UPDATE owner_table
	SET image_url = $1,
	updated_at = CURRENT_TIMESTAMP
	WHERE phone_number = $2`

	_, err := r.db.Exec(query, req.NewImageUrl, req.PhoneNumber)
	if err != nil {
		return &pba.UpdateOwnerImageResponse{
			Success: false,
		}, err
	}

	return &pba.UpdateOwnerImageResponse{
		Success: true,
	}, nil
}

// Get all owners by limit and pages
func (r *AttractionRepo) GetAllOwners(req *pba.GetAllOwnersRequest) (*pba.GetAllOwnersResponse, error) {
	// calculating offset
	offset := (req.Page - 1) * req.Limit

	query := `SELECT
	owner_id,
	full_name,
	email,
	password,
	birthday,
	phone_number,
	image_url,
	refresh_token,
	created_at,
	updated_at,
	deleted_at,
	role
	FROM owner_table
	LIMIT $1 OFFSET $2`

	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	var owners []*pba.Owner

	for rows.Next() {
		var owner pba.Owner

		err := rows.Scan(
			owner.OwnerId,
			owner.FullName,
			owner.Email,
			owner.Password,
			owner.Birthday,
			owner.PhoneNumber,
			owner.ImageUrl,
			owner.RefreshToken,
			owner.CreatedAt,
			owner.UpdatedAt,
			owner.DeletedAt,
			owner.Role,
		)
		if err != nil {
			// birorta owner ni to'ldirishda xatolik ketib qolsayam append qilib yuborishi uchun
			owners = append(owners, &owner)
		}

		owners = append(owners, &owner)
	}

	return &pba.GetAllOwnersResponse{
		Owners: owners,
	}, nil
}

// Create new attraction
func (r *AttractionRepo) CreateAttraction(attraction *pba.Attraction) (*pba.Attraction, error) {

	var (
		new_attraction pba.Attraction
		name, description, opening_hours, closing_hours, category,
		image_url, website_url, contact_information sql.NullString
	)
	query := `INSERT INTO attractions_table (
		attraction_id,
		name,
		description,
		location_id,
		opening_hours,
		closing_hours,
		category,
		rating,
		image_url,
		website_url,
		contact_information,
		owner_id)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING
		attraction_id,
		name,
		description,
		location_id,
		opening_hours,
		closing_hours,
		category,
		rating,
		image_url,
		website_url,
		contact_information,
		created_at,
		updated_at,
		owner_id`

	err := r.db.QueryRow(
		query,
		attraction.AttractionId,
		sql.NullString{String: attraction.Name, Valid: attraction.Name != ""},
		sql.NullString{String: attraction.Description, Valid: attraction.Description != ""},
		attraction.LocationId,
		sql.NullString{String: attraction.OpeningHours, Valid: attraction.OpeningHours != ""},
		sql.NullString{String: attraction.ClosingHours, Valid: attraction.ClosingHours != ""},
		sql.NullString{String: attraction.Category, Valid: attraction.Category != ""},
		attraction.Rating,
		sql.NullString{String: attraction.ImageUrl, Valid: attraction.ImageUrl != ""},
		sql.NullString{String: attraction.WebsiteUrl, Valid: attraction.WebsiteUrl != ""},
		sql.NullString{String: attraction.ContactInformation, Valid: attraction.ContactInformation != ""},
		attraction.OwnerId,
	).Scan(
		&new_attraction.AttractionId,
		&new_attraction.Name,
		&new_attraction.Description,
		&new_attraction.LocationId,
		&new_attraction.OpeningHours,
		&new_attraction.ClosingHours,
		&new_attraction.Category,
		&new_attraction.Rating,
		&new_attraction.ImageUrl,
		&new_attraction.WebsiteUrl,
		&new_attraction.ContactInformation,
		&new_attraction.CreatedAt,
		&new_attraction.UpdatedAt,
		&new_attraction.OwnerId,
	)

	if err != nil {
		return nil, err
	}

	// checking for sql null string
	new_attraction.Name = stringValue(name)
	new_attraction.Description = stringValue(description)
	new_attraction.OpeningHours = stringValue(opening_hours)
	new_attraction.ClosingHours = stringValue(closing_hours)
	new_attraction.Category = stringValue(category)
	new_attraction.ImageUrl = stringValue(image_url)
	new_attraction.WebsiteUrl = stringValue(website_url)
	new_attraction.ContactInformation = stringValue(contact_information)

	return &new_attraction, nil
}

// Get attraction(s) by name
func (r *AttractionRepo) GetAttractionByName(req *pba.GetAttractionByNameRequest) (*pba.GetAttractionByNameResponse, error) {
	query := `SELECT
	attraction_id,
	name,
	description,
	location_id,
	opening_hours,
	closing_hours,
	category,
	rating,
	image_url,
	website_url,
	contact_information,
	created_at,
	updated_at,
	owner_id
	FROM attractions_table
	WHERE name LIKE '%' || $1 || '%'`

	rows, err := r.db.Query(query, req.Name)
	if err != nil {
		return nil, err
	}

	var attractions []*pba.Attraction

	for rows.Next() {
		var attraction pba.Attraction

		err := rows.Scan(
			&attraction.AttractionId,
			&attraction.Name,
			&attraction.Description,
			&attraction.LocationId,
			&attraction.OpeningHours,
			&attraction.ClosingHours,
			&attraction.Category,
			&attraction.Rating,
			&attraction.ImageUrl,
			&attraction.WebsiteUrl,
			&attraction.ContactInformation,
			&attraction.CreatedAt,
			&attraction.UpdatedAt,
			&attraction.OwnerId,
		)
		if err != nil {
			attractions = append(attractions, &attraction)
		}

		attractions = append(attractions, &attraction)
	}

	return &pba.GetAttractionByNameResponse{
		Attractions: attractions,
	}, nil
}

// Get attraction(s) by category
func (r *AttractionRepo) GetAttractionsByCategory(req *pba.GetAttractionsByCategoryRequest) (*pba.GetAttractionsByCategoryResponse, error) {
	query := `SELECT
	attraction_id,
	name,
	description,
	location_id,
	opening_hours,
	closing_hours,
	category,
	rating,
	image_url,
	website_url,
	contact_information,
	created_at,
	updated_at,
	owner_id
	FROM attractions_table
	WHERE category = $1`

	rows, err := r.db.Query(query, req.Category)
	if err != nil {
		return nil, err
	}

	var attractions []*pba.Attraction

	for rows.Next() {
		var attraction pba.Attraction

		err := rows.Scan(
			&attraction.AttractionId,
			&attraction.Name,
			&attraction.Description,
			&attraction.LocationId,
			&attraction.OpeningHours,
			&attraction.ClosingHours,
			&attraction.Category,
			&attraction.Rating,
			&attraction.ImageUrl,
			&attraction.WebsiteUrl,
			&attraction.ContactInformation,
			&attraction.CreatedAt,
			&attraction.UpdatedAt,
			&attraction.OwnerId,
		)
		if err != nil {
			attractions = append(attractions, &attraction)
		}

		attractions = append(attractions, &attraction)
	}

	return &pba.GetAttractionsByCategoryResponse{
		Attractions: attractions,
	}, nil
}

// stringValue returns the string value of a sql.NullString, handling null values.
func stringValue(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
