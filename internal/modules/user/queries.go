package user

const (
	FindAllUsersQuery = `SELECT user_id, first_name, last_name, email, age FROM "user"`

	FindUserByIDQuery = `SELECT user_id, first_name, last_name, email, age FROM "user" WHERE user_id = $1`

	CreateUserQuery = `INSERT INTO users (first_name, last_name, email, age) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`
)
