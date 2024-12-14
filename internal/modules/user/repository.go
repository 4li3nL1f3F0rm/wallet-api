package user

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindAll() ([]User, error) {
	rows, err := r.DB.Query(FindAllUsersQuery)
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

func (r *UserRepository) FindById(id int) (User, error) {
	var user User
	err := r.DB.QueryRow(FindUserByIDQuery, id).
		Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return User{}, errors.New(fmt.Sprintf("user with id %d not found", id))
		}
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) Create(user User) (User, error) {
	err := r.DB.QueryRow(
		CreateUserQuery,
		user.FirstName, user.LastName, user.Email, user.Age,
	).Scan(&user.ID)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
