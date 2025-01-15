package queries

const (
	// Users CRUD Queries
	FetchUserByUsernameQuery = `
		SELECT id, username, password, access_level 
		FROM users 
		WHERE username = $1;
	`
	CreateUserQuery = `
		INSERT INTO users (id, username, password, access_level)
		VALUES ($1, $2, $3, $4);
	`
	UpdateUserQuery = `
		UPDATE users
		SET username = $2, password = $3, access_level = $4
		WHERE id = $1;
	`
	DeleteUserQuery = `
		DELETE FROM users
		WHERE id = $1;
	`

	// Houses CRUD Queries
	FetchHouseQuery = `
		SELECT id, street, builder, year_construction, created_at, update_at 
		FROM house 
		WHERE id = $1;
	`
	CreateHouseQuery = `
		INSERT INTO house (id, street, builder, year_construction, created_at, update_at)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	UpdateHouseQuery = `
		UPDATE house
		SET street = $2, builder = $3, year_construction = $4, update_at = $5
		WHERE id = $1;
	`
	DeleteHouseQuery = `
		DELETE FROM house
		WHERE id = $1;
	`

	FetchFlatQuery = `
		SELECT id, house_id, price, count_rooms, apartment_number, moderation_status 
		FROM flat
		WHERE id = $1;
	`
	CreateFlatQuery = `
		INSERT INTO flat (id, house_id, price, count_rooms, apartment_number, moderation_status)
		VALUES ($1, $2, $3, $4, $5, $6);
	`
	UpdateFlatQuery = `
		UPDATE flat
		SET house_id = $2, price = $3, count_rooms = $4, apartment_number = $5, moderation_status = $6
		WHERE id = $1;
	`
	DeleteFlatQuery = `
		DELETE FROM flat
		WHERE id = $1;
	`
)
