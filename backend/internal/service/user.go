package service

import (
	"superbank/internal/model"
	"superbank/internal/repository/postgres"
)


type CreateUserInput struct {
	Username string
	Password string 
	Role     string
}


type UserService interface {
	CreateUser(input model.CreateUserInput) error
	GetUserByUsername(username string) (*model.User, error)
}


type userService struct {
	userRepo postgres.UserRepository
}


func NewUserService(userRepo postgres.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}


func (s *userService) CreateUser(input model.CreateUserInput) error {
	return s.userRepo.CreateUser(input.Username, input.Password)
}


func (s *userService) GetUserByUsername(username string) (*model.User, error) {
	return s.userRepo.GetUserByUsername(username)
}
