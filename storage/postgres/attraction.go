package postgres

import (
	pba "attraction_service_booking/genproto/attraction_proto"
	"database/sql"
	"log"

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

//											**** OWNER ****

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

// Check uniqueness of upcoming owner account
func (r *AttractionRepo) CheckUniqueness(req *pba.UniquenessRequest) (*pba.UniquenessResponse, error) {
	return nil, nil
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

//											**** ATTRACTION ****

// Create a new attraction with image(s)
func (r *AttractionRepo) CreateAttraction(attraction *pba.Attraction) (*pba.Attraction, error) {

	// variables
	var (
		new_attraction pba.Attraction // response
		name, description, opening_hours, closing_hours, category,
		website_url, contact_information sql.NullString // prevent from empty value
		location_name, country, city, state_province, address sql.NullString // prevent from empty value
		images                                                []*pba.Image   // images response
		location                                              pba.Location   // location response
	)

	// query for locations_table
	query := `INSERT INTO locations_table (
		location_name,
		latitude,
		longitude,
		country,
		city,
		state_province,
		address,
		attraction_id)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING
		location_id,
		location_name,
		latitude,
		longitude,
		country,
		city,
		state_province,
		address,
		attraction_id`

	// insert location info into locations_table
	err := r.db.QueryRow(
		query,
		sql.NullString{String: attraction.Location.LocationName, Valid: attraction.Location.LocationName != ""},
		attraction.Location.Latitude,
		attraction.Location.Longitude,
		sql.NullString{String: attraction.Location.Country, Valid: attraction.Location.Country != ""},
		sql.NullString{String: attraction.Location.City, Valid: attraction.Location.City != ""},
		sql.NullString{String: attraction.Location.StateProvince, Valid: attraction.Location.StateProvince != ""},
		sql.NullString{String: attraction.Location.Address, Valid: attraction.Location.Address != ""},
		attraction.AttractionId,
	).Scan(
		&location.LocationId,
		&location.LocationName,
		&location.Latitude,
		&location.Longitude,
		&location.Country,
		&location.City,
		&location.StateProvince,
		&location.Address,
		&location.AttractionId,
	)
	if err != nil {
		return nil, err
	}

	// checking for sql null string
	attraction.Location.LocationName = stringValue(location_name)
	attraction.Location.Country = stringValue(country)
	attraction.Location.City = stringValue(city)
	attraction.Location.StateProvince = stringValue(state_province)
	attraction.Location.Address = stringValue(address)

	// insert images to images_table
	for _, image := range attraction.Images {
		var (
			attraction_id, image_url sql.NullString // prevent from empty data
			oneImage                 pba.Image      // one image response
		)

		query := `INSERT INTO images_table (
				attraction_id,
				image_url)
				VALUES($1, $2)
				RETURNING
				image_id,
				attraction_id,
				image_url,
				created_at`

		err := r.db.QueryRow(
			query,
			sql.NullString{String: image.AttractionId, Valid: image.AttractionId != ""},
			sql.NullString{String: image.ImageUrl, Valid: image.ImageUrl != ""},
		).Scan(
			&oneImage.ImageId,
			&oneImage.AttractionId,
			&oneImage.ImageUrl,
			&oneImage.CreatedAt,
		)

		if err != nil {
			// it is wrong that return nil because of inserting just an image into images_table
			log.Fatalf("Error while inserting images to images_table: %s", err)
		}

		// checking for empty string
		image.AttractionId = stringValue(attraction_id)
		image.ImageUrl = stringValue(image_url)

		// append one image to images array
		images = append(images, &oneImage)
	}

	// main query for attractions_table
	query = `INSERT INTO attractions_table (
		attraction_id,
		name,
		description,
		opening_hours,
		closing_hours,
		category,
		rating,
		website_url,
		contact_information,
		owner_id)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING
		attraction_id,
		name,
		description,
		opening_hours,
		closing_hours,
		category,
		rating,
		website_url,
		contact_information,
		created_at,
		updated_at,
		owner_id`

	err = r.db.QueryRow(
		query,
		attraction.AttractionId,
		sql.NullString{String: attraction.Name, Valid: attraction.Name != ""},
		sql.NullString{String: attraction.Description, Valid: attraction.Description != ""},
		sql.NullString{String: attraction.OpeningHours, Valid: attraction.OpeningHours != ""},
		sql.NullString{String: attraction.ClosingHours, Valid: attraction.ClosingHours != ""},
		sql.NullString{String: attraction.Category, Valid: attraction.Category != ""},
		attraction.Rating,
		sql.NullString{String: attraction.WebsiteUrl, Valid: attraction.WebsiteUrl != ""},
		sql.NullString{String: attraction.ContactInformation, Valid: attraction.ContactInformation != ""},
		attraction.OwnerId,
	).Scan(
		&new_attraction.AttractionId,
		&new_attraction.Name,
		&new_attraction.Description,
		&new_attraction.OpeningHours,
		&new_attraction.ClosingHours,
		&new_attraction.Category,
		&new_attraction.Rating,
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
	new_attraction.WebsiteUrl = stringValue(website_url)
	new_attraction.ContactInformation = stringValue(contact_information)

	// add images to response
	new_attraction.Images = images

	// add location info to response
	new_attraction.Location = &location

	return &new_attraction, nil
}

// Get attraction(s) by name
func (r *AttractionRepo) GetAttractionsByName(req *pba.GetAttractionByNameRequest) (*pba.GetAttractionByNameResponse, error) {
	query := `SELECT
	attraction_id,
	name,
	description,
	opening_hours,
	closing_hours,
	category,
	rating,
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
		var (
			attraction pba.Attraction
			location   pba.Location
		)

		err := rows.Scan(
			&attraction.AttractionId,
			&attraction.Name,
			&attraction.Description,
			&attraction.OpeningHours,
			&attraction.ClosingHours,
			&attraction.Category,
			&attraction.Rating,
			&attraction.WebsiteUrl,
			&attraction.ContactInformation,
			&attraction.CreatedAt,
			&attraction.UpdatedAt,
			&attraction.OwnerId,
		)
		if err != nil {
			return nil, err
		}

		// get location info query
		query := `SELECT
			location_id,
			location_name,
			latitude,
			longitude,
			country,
			city,
			state_province,
			address,
			attraction_id
			FROM locations_table
			WHERE attraction_id = $1`

		err = r.db.QueryRow(query, attraction.AttractionId).Scan(
			&location.LocationId,
			&location.LocationName,
			&location.Latitude,
			&location.Longitude,
			&location.Country,
			&location.City,
			&location.StateProvince,
			&location.Address,
			&location.AttractionId,
		)
		if err != nil {
			log.Fatalf("Error while getting location info: %s", err)
		}

		// add location info to response attraction
		attraction.Location = &location

		query = `SELECT
			image_id,
			attraction_id,
			image_url,
			created_at
			FROM images_table
			WHERE attraction_id = $1`

		rows, err := r.db.Query(query, attraction.AttractionId)
		if err != nil {
			log.Fatalf("Error while getting images: %s", err)
		}

		var images []*pba.Image

		for rows.Next() {
			var image pba.Image

			err := rows.Scan(
				&image.ImageId,
				&image.AttractionId,
				&image.ImageUrl,
				&image.CreatedAt,
			)
			if err != nil {
				log.Fatalf("Error while getting an image: %s", err)
			}

			// add an image to response images
			images = append(images, &image)
		}

		// add images info to attraction response
		attraction.Images = images

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
	opening_hours,
	closing_hours,
	category,
	rating,
	website_url,
	contact_information,
	created_at,
	updated_at,
	owner_id
	FROM attractions_table
	WHERE category LIKE '%' || $1 || '%'`

	rows, err := r.db.Query(query, req.Category)
	if err != nil {
		return nil, err
	}

	var attractions []*pba.Attraction

	for rows.Next() {
		var (
			attraction pba.Attraction
			location   pba.Location
		)

		err := rows.Scan(
			&attraction.AttractionId,
			&attraction.Name,
			&attraction.Description,
			&attraction.OpeningHours,
			&attraction.ClosingHours,
			&attraction.Category,
			&attraction.Rating,
			&attraction.WebsiteUrl,
			&attraction.ContactInformation,
			&attraction.CreatedAt,
			&attraction.UpdatedAt,
			&attraction.OwnerId,
		)
		if err != nil {
			return nil, err
		}

		// get location info query
		query := `SELECT
			location_id,
			location_name,
			latitude,
			longitude,
			country,
			city,
			state_province,
			address,
			attraction_id
			FROM locations_table
			WHERE attraction_id = $1`

		err = r.db.QueryRow(query, attraction.AttractionId).Scan(
			&location.LocationId,
			&location.LocationName,
			&location.Latitude,
			&location.Longitude,
			&location.Country,
			&location.City,
			&location.StateProvince,
			&location.Address,
			&location.AttractionId,
		)
		if err != nil {
			log.Fatalf("Error while getting location info: %s", err)
		}

		// add location info to response attraction
		attraction.Location = &location

		query = `SELECT
			image_id,
			attraction_id,
			image_url,
			created_at
			FROM images_table
			WHERE attraction_id = $1`

		rows, err := r.db.Query(query, attraction.AttractionId)
		if err != nil {
			log.Fatalf("Error while getting images: %s", err)
		}

		var images []*pba.Image

		for rows.Next() {
			var image pba.Image

			err := rows.Scan(
				&image.ImageId,
				&image.AttractionId,
				&image.ImageUrl,
				&image.CreatedAt,
			)
			if err != nil {
				log.Fatalf("Error while getting an image: %s", err)
			}

			// add an image to response images
			images = append(images, &image)
		}

		// add images info to attraction response
		attraction.Images = images

		attractions = append(attractions, &attraction)
	}

	return &pba.GetAttractionsByCategoryResponse{
		Attractions: attractions,
	}, nil
}

// Get attraction(s) by location
func (r *AttractionRepo) GetAttractionsByLocation(req *pba.GetAttractionsByLocationRequest) (*pba.GetAttractionsByLocationResponse, error) {
	// it is for response attractions
	var attractions []*pba.Attraction

	// query: get all attraction id where located in the city
	query := `SELECT attraction_id FROM locations_table WHERE city = $1`

	// several location's attraction_id
	rows, err := r.db.Query(query, req.Location.City)
	if err != nil {
		return nil, nil
	}

	for rows.Next() {

		var attraction_id string

		err := rows.Scan(&attraction_id)
		if err != nil {
			return nil, err
		}

		// query: get all attractions by attraction_id
		query := `SELECT
		attraction_id,
		name,
		description,
		opening_hours,
		closing_hours,
		category,
		rating,
		website_url,
		contact_information,
		created_at,
		updated_at,
		owner_id
		FROM attractions_table
		WHERE attraction_id = $1`

		// several attractions
		rows, err := r.db.Query(query, attraction_id)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var (
				attraction pba.Attraction
				location   pba.Location
			)

			err := rows.Scan(
				&attraction.AttractionId,
				&attraction.Name,
				&attraction.Description,
				&attraction.OpeningHours,
				&attraction.ClosingHours,
				&attraction.Category,
				&attraction.Rating,
				&attraction.WebsiteUrl,
				&attraction.ContactInformation,
				&attraction.CreatedAt,
				&attraction.UpdatedAt,
				&attraction.OwnerId,
			)
			if err != nil {
				return nil, err
			}

			// query: get location info query
			query := `SELECT
				location_id,
				location_name,
				latitude,
				longitude,
				country,
				city,
				state_province,
				address,
				attraction_id
				FROM locations_table
				WHERE attraction_id = $1`

			err = r.db.QueryRow(query, attraction.AttractionId).Scan(
				&location.LocationId,
				&location.LocationName,
				&location.Latitude,
				&location.Longitude,
				&location.Country,
				&location.City,
				&location.StateProvince,
				&location.Address,
				&location.AttractionId,
			)
			if err != nil {
				log.Fatalf("Error while getting location info: %s", err)
			}

			// add location info to response attraction
			attraction.Location = &location

			// query: get all images of attraction
			query = `SELECT
				image_id,
				attraction_id,
				image_url,
				created_at
				FROM images_table
				WHERE attraction_id = $1`

			rows, err := r.db.Query(query, attraction.AttractionId)
			if err != nil {
				log.Fatalf("Error while getting images: %s", err)
			}

			var images []*pba.Image

			for rows.Next() {
				var image pba.Image

				err := rows.Scan(
					&image.ImageId,
					&image.AttractionId,
					&image.ImageUrl,
					&image.CreatedAt,
				)
				if err != nil {
					log.Fatalf("Error while getting an image: %s", err)
				}

				// add an image to response images
				images = append(images, &image)
			}

			// add images info to attraction response
			attraction.Images = images

			attractions = append(attractions, &attraction)
		}
	}

	return &pba.GetAttractionsByLocationResponse{
		Attractions: attractions,
	}, nil
}

// Get attraction(s) by rating
func (r *AttractionRepo) GetAttractionsByRating(req *pba.GetAttractionsByRatingRequest) (*pba.GetAttractionsByRatingResponse, error) {

	// calculating offset
	offset := (req.Page - 1) * req.Limit

	// query: get attractions by rating from higher to lower by using LIMIT and OFFSET
	query := `SELECT
	attraction_id,
	name,
	description,
	opening_hours,
	closing_hours,
	category,
	rating,
	website_url,
	contact_information,
	created_at,
	updated_at,
	owner_id
	FROM attractions_table
	LIMIT $1 OFFSET $2
	ORDER BY rating DESC`

	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}

	var attractions []*pba.Attraction

	for rows.Next() {
		var (
			attraction pba.Attraction
			location   pba.Location
		)

		err := rows.Scan(
			&attraction.AttractionId,
			&attraction.Name,
			&attraction.Description,
			&attraction.OpeningHours,
			&attraction.ClosingHours,
			&attraction.Category,
			&attraction.Rating,
			&attraction.WebsiteUrl,
			&attraction.ContactInformation,
			&attraction.CreatedAt,
			&attraction.UpdatedAt,
			&attraction.OwnerId,
		)
		if err != nil {
			return nil, err
		}

		// get location info query
		query := `SELECT
			location_id,
			location_name,
			latitude,
			longitude,
			country,
			city,
			state_province,
			address,
			attraction_id
			FROM locations_table
			WHERE attraction_id = $1`

		err = r.db.QueryRow(query, attraction.AttractionId).Scan(
			&location.LocationId,
			&location.LocationName,
			&location.Latitude,
			&location.Longitude,
			&location.Country,
			&location.City,
			&location.StateProvince,
			&location.Address,
			&location.AttractionId,
		)
		if err != nil {
			log.Fatalf("Error while getting location info: %s", err)
		}

		// add location info to response attraction
		attraction.Location = &location

		query = `SELECT
			image_id,
			attraction_id,
			image_url,
			created_at
			FROM images_table
			WHERE attraction_id = $1`

		rows, err := r.db.Query(query, attraction.AttractionId)
		if err != nil {
			log.Fatalf("Error while getting images: %s", err)
		}

		var images []*pba.Image

		for rows.Next() {
			var image pba.Image

			err := rows.Scan(
				&image.ImageId,
				&image.AttractionId,
				&image.ImageUrl,
				&image.CreatedAt,
			)
			if err != nil {
				log.Fatalf("Error while getting an image: %s", err)
			}

			// add an image to response images
			images = append(images, &image)
		}

		// add images info to attraction response
		attraction.Images = images

		attractions = append(attractions, &attraction)
	}

	return &pba.GetAttractionsByRatingResponse{
		Attractions: attractions,
	}, nil
}

// Update attraction image by attraction_id
func (r *AttractionRepo) UpdateAttractionImage(req *pba.UpdateAttractionImageRequest) (*pba.UpdateAttractionImageResponse, error) {

	// delete previous images
	query := `DELETE FROM images_table WHERE attraction_id = $1`
	_, err := r.db.Exec(query, req.AttractionId)
	if err != nil {
		return nil, err
	}

	// getting each image from request images
	for _, image := range req.Images {
		query := `INSERT INTO images_table (
			attraction_id,
			image_url)
			VALUES($1, $2)`

		_, err := r.db.Exec(query, image.AttractionId, image.ImageUrl)
		if err != nil {
			return &pba.UpdateAttractionImageResponse{
				Success: false,
			}, err
		}
	}

	return &pba.UpdateAttractionImageResponse{
		Success: true,
	}, nil
}

// Delete attraction image(s)
func (r *AttractionRepo) DeleteAttractionImage(req *pba.DeleteAttractionImageRequest) (*pba.DeleteAttractionImageResponse, error) {
	// query: set image(s) of attraction to an empty string by id
	query := `UPDATE images_table
			SET image_url = ''
			WHERE attraction_id = $1`

	_, err := r.db.Exec(query, req.AttractionId)
	if err != nil {
		return &pba.DeleteAttractionImageResponse{
			Success: false,
		}, err
	}

	return &pba.DeleteAttractionImageResponse{
		Success: true,
	}, nil
}

//											**** REVIEW ****

// add a new review
func (r *AttractionRepo) AddReview(req *pba.Review) (*pba.Review, error) {
	var (
		review  pba.Review
		comment sql.NullString
	)

	// query: add a new query to reviews_table
	qeury := `INSERT INTO reviews_table (
		review_id,
		attraction_id,
		user_id,
		rating,
		comment)
		VALUES($1, $2, $3, $4, $5)
		RETURNING
		review_id,
		attraction_id,
		user_id,
		rating,
		comment,
		created_at`

	err := r.db.QueryRow(
		qeury,
		req.ReviewId,
		req.AttractionId,
		req.UserId,
		req.Rating,
		sql.NullString{String: req.Comment, Valid: req.Comment != ""},
	).Scan(
		&review.ReviewId,
		&review.AttractionId,
		&review.UserId,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	// check for null string
	req.Comment = stringValue(comment)

	// returning response
	return &review, nil
}

// get review by id
func (r *AttractionRepo) GetReview(req *pba.GetReviewRequest) (*pba.GetReviewResponse, error) {
	var review pba.Review
	// query: get a review by id
	query := `SELECT review_id, attraction_id, user_id, rating, comment, created_at FROM reviews_table WHERE review_id = $1`

	err := r.db.QueryRow(query, req.ReviewId).Scan(
		&review.ReviewId,
		&review.AttractionId,
		&review.UserId,
		&review.Rating,
		&review.Comment,
		&review.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &pba.GetReviewResponse{
		Review: &review,
	}, nil
}

// list of reviews of one attraction
func (r *AttractionRepo) ListReviews(req *pba.ListReviewsRequest) (*pba.ListReviewsResponse, error) {
	var reviews []*pba.Review

	// query: get list of reviews by attraction_id
	query := `SELECT review_id, attraction_id, user_id, rating, comment, created_at FROM reviews_table WHERE attraction_id = $1`

	rows, err := r.db.Query(query, req.AttractionId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var review pba.Review

		err := rows.Scan(
			&review.ReviewId,
			&review.AttractionId,
			&review.UserId,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
		)
		if err != nil {
			log.Fatalf("Error while getting a review: %s", err)
		}

		// add a review to reviews array
		reviews = append(reviews, &review)
	}

	return &pba.ListReviewsResponse{
		Reviews: reviews,
	}, nil
}

// update review comment by id
func (r *AttractionRepo) UpdateReviewComment(req *pba.UpdateReviewCommentRequest) (*pba.UpdateReviewCommentResponse, error) {
	// query: update review comment by review_id
	query := `UPDATE reviews_table
			SET comment = $1
			WHERE review_id = $2`

	_, err := r.db.Exec(query, req.NewComment, req.ReviewId)
	if err != nil {
		return &pba.UpdateReviewCommentResponse{
			Success: false,
		}, err
	}

	return &pba.UpdateReviewCommentResponse{
		Success: true,
	}, nil
}

// delete review by id
func (r *AttractionRepo) DeleteReview(req *pba.DeleteReviewRequest) (*pba.DeleteReviewResponse, error) {
	// query: delete a review by review_id
	query := `DELETE FROM reviews_table WHERE review_id = $1`

	_, err := r.db.Exec(query, req.ReviewId)
	if err != nil {
		return &pba.DeleteReviewResponse{
			Success: false,
		}, err
	}

	return &pba.DeleteReviewResponse{
		Success: true,
	}, nil
}

//											**** FAVOURITE ****

// add attraction to favourites_table
func (r *AttractionRepo) AddToFavourites(req *pba.AddToFavouritesRequest) (*pba.AddToFavouritesResponse, error) {
	var response pba.Favourite

	// query: insert into favourites_table
	query := `INSERT INTO favourites_table(
		favourite_id,
		user_id,
		attraction_id)
		VALUES($1, $2, $3)
		RETURNING
		favourite_id,
		user_id,
		attraction_id`

	err := r.db.QueryRow(
		query,
		req.Favourite.FavouriteId,
		req.Favourite.UserId,
		req.Favourite.AttractionId,
	).Scan(
		&response.FavouriteId,
		&response.UserId,
		&response.AttractionId,
	)
	if err != nil {
		return nil, err
	}

	return &pba.AddToFavouritesResponse{
		Favourite: &response,
	}, nil
}

// drop from favourites list
func (r *AttractionRepo) DropFromFavourites(req *pba.DropFromFavouritesRequest) (*pba.DropFromFavouritesResponse, error) {
	query := `DELETE FROM favourites_table WHERE id = $1`

	_, err := r.db.Exec(query, req.AttractionId)
	if err != nil {
		return &pba.DropFromFavouritesResponse{
			Success: false,
		}, err
	}

	return &pba.DropFromFavouritesResponse{
		Success: true,
	}, nil
}

// list of favourites by user_id
func (r *AttractionRepo) ListOfFavourites(req *pba.ListOfFavouritesRequest) (*pba.ListOfFavouritesResponse, error) {
	var attractions []*pba.Attraction

	// query:
	query := `SELECT attraction_id FROM favourites_table WHERE user_id = $1`

	rows, err := r.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			attraction_id string
			attraction    pba.Attraction
		)

		// catch attraction_id
		err := rows.Scan(attraction_id)
		if err != nil {
			return nil, err
		}

		// query: get attraction by attraction_id
		query := `SELECT
		attraction_id,
		name,
		description,
		opening_hours,
		closing_hours,
		category,
		rating,
		website_url,
		contact_information,
		created_at,
		updated_at,
		owner_id
		FROM attractions_table
		WHERE attraction_id = $1`

		err = r.db.QueryRow(query, attraction_id).Scan(
			&attraction.AttractionId,
			&attraction.Name,
			&attraction.Description,
			&attraction.OpeningHours,
			&attraction.ClosingHours,
			&attraction.Rating,
			&attraction.WebsiteUrl,
			&attraction.ContactInformation,
			&attraction.CreatedAt,
			&attraction.UpdatedAt,
			&attraction.OwnerId,
		)
		if err != nil {
			return nil, err
		}

		attractions = append(attractions, &attraction)
	}

	return &pba.ListOfFavouritesResponse{
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
