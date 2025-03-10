package service

import (
	"time"

	"superbank/internal/model"
	"superbank/internal/repository/postgres"
	"superbank/pkg/util"
)

type AuthService interface {
	Login(req model.LoginRequest) (string, error)
	ValidateToken(tokenString string) (string, string, error)
}

type authService struct {
	userRepo  postgres.UserRepository
	jwtSecret string
}

func NewAuthService(userRepo postgres.UserRepository, jwtSecret string) AuthService {
	return &authService{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (s *authService) Login(req model.LoginRequest) (string, error) {
	
	user, err := s.userRepo.ValidateCredentials(req.Username, req.Password)
	if err != nil {
		return "", err
	}

	
	userID := string(user.ID) 
	return util.GenerateJWT(s.jwtSecret, userID, user.Role, 24*time.Hour)
}

func (s *authService) ValidateToken(tokenString string) (string, string, error) {
	return util.ValidateJWT(s.jwtSecret, tokenString)
}
