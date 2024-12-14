package user

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.DB.Query(`SELECT user_id, first_name, last_name, email, age FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Create(user User) (User, error) {
	err := r.DB.QueryRow(
		"INSERT INTO users (first_name, last_name, email, age) VALUES ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Email, user.Age,
	).Scan(&user.ID)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
