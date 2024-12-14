package user

import "errors"

type UserService struct {
	Repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) GetUser(id int) (User, error) {
	return s.Repo.FindById(id)
}

func (s *UserService) CreateUser(input CreateUserRequest) (User, error) {
	if input.Email == "" {
		return User{}, errors.New("email is required")
	}

	user := User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Age:       input.Age,
	}
	return s.Repo.Create(user)
}
